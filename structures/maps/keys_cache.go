package main

import "fmt"

type Person struct {
	Name *string
	Age  *int
}

func main() {
	m := make(map[*int]struct{})
	m[new(int)] = struct{}{}
	m[new(int)] = struct{}{}
	m[new(int)] = struct{}{}
	m[new(int)] = struct{}{}
	fmt.Println(m)

	m2 := make(map[chan string]struct{})
	m2[make(chan string)] = struct{}{}
	fmt.Println(m2)
}
