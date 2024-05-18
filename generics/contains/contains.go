package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 4, 8}

	fmt.Println(IsContains(list, 3))
	fmt.Println(IsContains(list, 4))
}

func IsContains[T comparable](list []T, n T) bool {
	for _, i := range list {
		if i == n {
			return true
		}
	}
	return false
}
