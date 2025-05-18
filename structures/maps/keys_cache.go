package main

import "fmt"

type Person struct {
	Name *string
	Age  *int
}

type UnsafeStruct struct {
	Data []int       // слайс - несравнимый тип
	Map  map[int]int // map - несравнимый тип
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

	// Это НЕ будет работать
	m3 := make(map[UnsafeStruct]string)

}
