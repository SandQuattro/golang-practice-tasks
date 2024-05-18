package main

import "fmt"

type P struct {
	A int
}

func main() {
	arr := []P{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}}

	for _, v := range arr {
		if v.A == 4 {
			arr = append(arr, P{A: 10})
		}
		fmt.Println(v.A)
	}
}
