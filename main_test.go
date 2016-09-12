package main

import (
	"fmt"
	"testing"
)

func BenchmarkCollisionFourBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val1, val2, hash := generateBirthdayCollision(4)
		fmt.Printf("%v, %v, %x", val1, val2, hash)
	}
}

//
// func BenchmarkCollisionSixBits(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 	}
// }
// func BenchmarkCollisionEightBits(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 	}
// }
// func BenchmarkCollisionTenBits(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 	}
// }
// func BenchmarkCollisionElevenBits(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 	}
// }
// func BenchmarkCollisionTwelveBits(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 	}
// }
//
// func BenchmarkCollisionThirteenBits(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 	}
// }
//
// func BenchmarkCollisionFourteenBits(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 	}
// }
// func BenchmarkCollisionFifteenBits(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 	}
// }
// func BenchmarkCollisionSixteenTwoBits(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 	}
// }
