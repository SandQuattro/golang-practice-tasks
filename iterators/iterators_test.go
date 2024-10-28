package main

import (
	"testing"
)

var (
	transform = func(i int) int { return i * 2 }
	list      = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

// In this case, the iterator is much faster! Why? Because the iterator doesn’t transform the entire list—it stops
// as soon as it finds the result you’re looking for. On the other hand, NormalTransform still transforms
// the entire list, even if we only care about one item.
func BenchmarkNormalTransform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, num := range NormalTransform(list, transform) {
			if num == 4 {
				break
			}
		}
	}
}

func BenchmarkIteratorTransform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, num := range IteratorTransform(list, transform) {
			if num == 4 {
				break
			}
		}
	}
}
