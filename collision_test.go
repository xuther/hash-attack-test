package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func BenchmarkCollisionFourBits(b *testing.B) {
	f, err := os.Create("4Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(4)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))
	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}

// func TestCollisionFourBits(b *testing.T) {
// 	val1, val2, hash := generateBirthdayCollision(4)
// 	//fmt.Printf("%v, %v, %x", val1, val2, hash)
// }

func BenchmarkCollisionSixBits(b *testing.B) {
	f, err := os.Create("6Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(6)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}

func BenchmarkCollisionEightBits(b *testing.B) {

	f, err := os.Create("8Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(8)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionTenBits(b *testing.B) {

	f, err := os.Create("10Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(10)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionElevenBits(b *testing.B) {

	f, err := os.Create("11Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(11)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionTwelveBits(b *testing.B) {

	f, err := os.Create("12Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(12)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}

func BenchmarkCollisionThirteenBits(b *testing.B) {

	f, err := os.Create("13Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(13)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}

func BenchmarkCollisionFourteenBits(b *testing.B) {

	f, err := os.Create("14Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(14)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionFifteenBits(b *testing.B) {

	f, err := os.Create("15Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(15)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionSixteenBits(b *testing.B) {

	f, err := os.Create("16Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(16)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionSeventeenBits(b *testing.B) {

	f, err := os.Create("17Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(17)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionEighteenBits(b *testing.B) {

	f, err := os.Create("18Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(18)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionNineteenBits(b *testing.B) {

	f, err := os.Create("19Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(19)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionTwentyBits(b *testing.B) {

	f, err := os.Create("20Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(20)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionTwentyOneBits(b *testing.B) {

	f, err := os.Create("21Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(21)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionTwentyTwoBits(b *testing.B) {

	f, err := os.Create("22Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(22)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionTwentyThreeBits(b *testing.B) {

	f, err := os.Create("23Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(23)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionTwentyFourBits(b *testing.B) {

	f, err := os.Create("24Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(24)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionTwentyEightBits(b *testing.B) {

	f, err := os.Create("28Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(28)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionThirtyTwoBits(b *testing.B) {

	f, err := os.Create("32Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(32)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionThirtySixBits(b *testing.B) {

	f, err := os.Create("36Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(36)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionFourtyBits(b *testing.B) {

	f, err := os.Create("40Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(40)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionFourtyFourBits(b *testing.B) {

	f, err := os.Create("44Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(44)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkCollisionFourtyEightBits(b *testing.B) {

	f, err := os.Create("48Collision.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generateBirthdayCollision(48)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
