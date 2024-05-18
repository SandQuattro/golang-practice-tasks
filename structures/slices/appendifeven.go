package main

import (
	"fmt"
)

func modify(s []int) {
	for i, n := range s {
		s[i] = n * 2
		if i%2 == 0 {
			s = append(s, i*2)
		}
	}
}

func main() {
	sl := []int{1, 2, 3}
	modify(sl)
	fmt.Println(sl)
}
