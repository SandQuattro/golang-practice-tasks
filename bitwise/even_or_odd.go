package main

import "fmt"

func main() {
	fmt.Println(isEven(0))
	fmt.Println(isEven(1))
	fmt.Println(isEven(2))
	fmt.Println(isEven(3))

	fmt.Println(isEvenFast(0))
	fmt.Println(isEvenFast(1))
	fmt.Println(isEvenFast(2))
	fmt.Println(isEvenFast(3))
}

func isEven(num int) bool {
	return num%2 == 0
}

func isEvenFast(num int) bool {
	return num&1 == 0
}
