package main

import (
	"testing"
)

const size = 10000

func sliceIterator(sl []int) {
	var cnt int
	for i := 0; i < size; i++ {
		cnt += sl[i]
	}
}

func mapIterator(m map[int]int) {
	var cnt int
	for i := 0; i < size; i++ {
		cnt += m[i]
	}
}

func BenchmarkSliceReading(b *testing.B) {
	sl := make([]int, size)
	for i := 0; i < size; i++ {
		sl[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sliceIterator(sl)
	}
}

func BenchmarkMapReading(b *testing.B) {
	m := make(map[int]int, size)
	for i := 0; i < size; i++ {
		m[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mapIterator(m)
	}
}
