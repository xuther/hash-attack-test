package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%^&*()"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func generateBirthdayCollision(length int) (string, string, string) {
	triedHashes := make(map[string]string)

	for true {
		strVal := getRandomString(55)
		hash := string(shaWrapper(strVal, length))
		if val, ok := triedHashes[hash]; ok {
			return val, strVal, hash
		}

		triedHashes[hash] = strVal
	}
	return "", "", ""
}

func getRandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

//length is in bits
func shaWrapper(toHash string, length int) []byte {

	bitsPreProc := sha1.Sum([]byte(toHash))

	remain := length % 8
	fmt.Printf("length: %v, remain: %v\n", length, remain)
	if remain == 0 {
		return bitsPreProc[:(length / 8)]
	}

	//now we need to mask the number of
	bitsPostProc := bitsPreProc[:(length/8)+1]
	fmt.Printf("Bits: %s\n", string(bitsPostProc))
	//Could do something other than switch (some sort of loop) but this seems simpler
	switch remain {
	case 1:
		bitsPostProc[len(bitsPostProc)-1] &= 0x01
		break
	case 2:
		bitsPostProc[len(bitsPostProc)-1] &= 0x03
		break
	case 3:
		bitsPostProc[len(bitsPostProc)-1] &= 0x07
		break
	case 4:
		bitsPostProc[len(bitsPostProc)-1] &= 0x0f
		break
	case 5:
		bitsPostProc[len(bitsPostProc)-1] &= 0x1f
		break
	case 6:
		bitsPostProc[len(bitsPostProc)-1] &= 0x3f
		break
	case 7:
		bitsPostProc[len(bitsPostProc)-1] &= 0x7f
		break
	}
	fmt.Printf("Bits: %+v\n", bitsPostProc)
	return bitsPostProc
}

func main() {
	fmt.Printf("Test")
	val1, val2, hash := generateBirthdayCollision(4)
	fmt.Printf("%v, %v, %x", val1, val2, hash)
}
