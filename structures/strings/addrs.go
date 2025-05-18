package main

import "fmt"

func main() {
	a := "hello"
	fmt.Println("address of a is ", &a)
	a = "world"
	fmt.Println("address of a is ", &a)
}
