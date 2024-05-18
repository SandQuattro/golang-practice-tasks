package main

import "fmt"

func main() {
	//  Slices in Go are always passed by reference.

	s1 := []int{1, 5, 7, 12, 9}
	reslice := s1[1:3] // [5 7]

	fmt.Println(s1)      // [1 5 7 12 9]
	fmt.Println(reslice) // [5 7]

	// make some changes
	reslice[0] = 128
	reslice[1] = 256

	fmt.Println(s1)      // [1 128 256 12 9]
	fmt.Println(reslice) // [128 256]
}
