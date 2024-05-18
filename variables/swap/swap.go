package main

import "fmt"

func main() {
	fmt.Println(swap(15, 10))
}

func swap(a, b int) (int, int) {
	fmt.Println(a, b)
	b, a = a, b
	return a, b
}
