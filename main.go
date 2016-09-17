package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%^&*()"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func generatePreImageAttack(length int) (string, string, string, int) {
	OrigVal := getRandomString(55)
	Orighash := string(shaWrapper(OrigVal, length))
	count := 0

	for true {
		count = count + 1
		strVal := getRandomString(55)
		hash := string(shaWrapper(strVal, length))
		if hash == Orighash {
			return OrigVal, strVal, hash, count
		}
	}
	return "", "", "", count
}

func generateBirthdayCollision(length int) (string, string, string, int) {
	triedHashes := make(map[string]string)
	count := 0

	for true {
		count = count + 1
		strVal := getRandomString(55)
		hash := string(shaWrapper(strVal, length))
		if val, ok := triedHashes[hash]; ok {
			return val, strVal, hash, count
		}

		triedHashes[hash] = strVal
	}
	return "", "", "", count
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
	if remain == 0 {
		return bitsPreProc[:(length / 8)]
	}

	//now we need to mask the number of
	bitsPostProc := bitsPreProc[:(length/8)+1]
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
	return bitsPostProc
}

func avgAttempts(values []int) float64 {
	running := 0
	for i := 0; i < len(values); i++ {
		running = running + values[i]
	}

	return float64(running) / float64(len(values))
}

func main() {

	f, err := os.Create("deviation.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()
	N := 10000
	fmt.Fprintf(f, "%v\t%v\t%v\t%v\t%v\t%v\n", "attempts", "Length", "Hash1", "Hash2", "String1", "String2")

	for i := 0; i < N; i++ {
		preTimestamp := time.Now()
		_, _, _, attempt := generateBirthdayCollision(24)
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v", N, avgAttempts(attempts))
}
