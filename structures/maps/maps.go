package main

import "fmt"

func main() {
	m := make(map[int]int, 1000)
	elem := &m[1]
	fmt.Sprintf("element 1 = %d", elem)
}
