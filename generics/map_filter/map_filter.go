package main

import "fmt"

func main() {
	list := []int{1, 2, 4, 8}
	inc := func(i int) int { return i + 1 }

	res := Map(list, inc)
	fmt.Println(res)

	list = []int{1, 2, 4, 8}
	filter := func(i int) bool {
		return i == 4
	}
	res = Filter(list, filter)
	fmt.Println(res)
}

func Map[T any](list []T, f func(T) T) []T {
	for i, el := range list {
		list[i] = f(el)
	}
	return list
}

func Filter[T any](list []T, f func(T) bool) []T {
	var res []T
	for _, el := range list {
		if f(el) {
			res = append(res, el)
		}
	}
	return res
}
