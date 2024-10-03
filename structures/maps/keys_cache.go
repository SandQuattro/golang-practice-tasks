package main

import "fmt"

func main() {
	m := make(map[*int]struct{})
	m[new(int)] = struct{}{}
	m[new(int)] = struct{}{}
	m[new(int)] = struct{}{}
	m[new(int)] = struct{}{}
	fmt.Println(m)
}
