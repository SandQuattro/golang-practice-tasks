package main

import "fmt"

func main() {
	var stackVar int = 5
	var heapVar *int = new(int)

	fmt.Printf("stackVar: %d heapVar: %p", stackVar, heapVar)
}
