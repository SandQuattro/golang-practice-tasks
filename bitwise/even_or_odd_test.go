package main

import "testing"

func BenchmarkIsEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEven(i)
	}
}

func BenchmarkIsEvenFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEvenFast(i)
	}
}
