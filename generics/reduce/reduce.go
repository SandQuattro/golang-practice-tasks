package main

import "fmt"

func main() {
	list := []int{1, 2, 4, 8}

	sum := func(x, y int) int { return x + y }

	res := Reduce(list, sum, 0)
	fmt.Println(res)

	mult := func(x, y int) int { return x * y }
	res = Reduce(list, mult, 1)
	fmt.Println(res)
}

func Reduce[T any](list []T, accumulator func(T, T) T, initial T) T {
	for _, el := range list {
		// тут начальное значение или на след итерации тут будет результат предыдущего действия
		initial = accumulator(initial, el)
	}
	return initial
}
