package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func BenchmarkPreImageFourBits(b *testing.B) {
	f, err := os.Create("4PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(4)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))
	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}

// func TestPreImageFourBits(b *testing.T) {
// 	val1, val2, hash := generatePreImageAttack(4)
// 	//fmt.Printf("%v, %v, %x", val1, val2, hash)
// }

func BenchmarkPreImageSixBits(b *testing.B) {
	f, err := os.Create("6PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(6)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}

func BenchmarkPreImageEightBits(b *testing.B) {

	f, err := os.Create("8PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(8)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageTenBits(b *testing.B) {

	f, err := os.Create("10PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(10)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageElevenBits(b *testing.B) {

	f, err := os.Create("11PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(11)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageTwelveBits(b *testing.B) {

	f, err := os.Create("12PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(12)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}

func BenchmarkPreImageThirteenBits(b *testing.B) {

	f, err := os.Create("13PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(13)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}

func BenchmarkPreImageFourteenBits(b *testing.B) {

	f, err := os.Create("14PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(14)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageFifteenBits(b *testing.B) {

	f, err := os.Create("15PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(15)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageSixteenBits(b *testing.B) {

	f, err := os.Create("16PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(16)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageSeventeenBits(b *testing.B) {

	f, err := os.Create("17PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(17)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageEighteenBits(b *testing.B) {

	f, err := os.Create("18PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(18)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageNineteenBits(b *testing.B) {

	f, err := os.Create("19PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(19)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageTwentyBits(b *testing.B) {

	f, err := os.Create("20PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(20)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageTwentyOneBits(b *testing.B) {

	f, err := os.Create("21PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(21)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageTwentyTwoBits(b *testing.B) {

	f, err := os.Create("22PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(22)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageTwentyThreeBits(b *testing.B) {

	f, err := os.Create("23PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(23)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageTwentyFourBits(b *testing.B) {

	f, err := os.Create("24PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(24)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
func BenchmarkPreImageTwentyEightBits(b *testing.B) {

	f, err := os.Create("28PreImage.txt")
	if err != nil {
		panic(err)
	}
	attempts := []int{}

	defer f.Close()

	for i := 0; i < b.N; i++ {
		preTimestamp := time.Now()
		b.StartTimer()
		_, _, _, attempt := generatePreImageAttack(28)
		b.StopTimer()
		d := time.Since(preTimestamp)

		fmt.Fprintf(f, "%v\t%v\n", attempt, d.Nanoseconds())
		attempts = append(attempts, attempt)
	}

	fmt.Printf("%v \t %v\n", b.N, avgAttempts(attempts))

	//fmt.Printf("%v, %v, %x", val1, val2, hash)
}
