package main

import "fmt"

func main() {
	var x []int
	x = append(x, 0)
	fmt.Printf("len x:%d cap x:%d, pointer x:%p\n", len(x), cap(x), &x)
	x = append(x, 1)
	fmt.Printf("len x:%d cap x:%d, pointer x:%p\n", len(x), cap(x), &x)
	x = append(x, 2)
	fmt.Printf("len x:%d cap x:%d, pointer x:%p\n", len(x), cap(x), &x)

	y := append(x, 3)

	fmt.Printf("len x:%d cap x:%d, pointer x[0]:%p\n", len(x), cap(x), &x[2])
	fmt.Printf("len y:%d cap y:%d, pointer y[0]:%p\n", len(y), cap(y), &y[2])

	fmt.Println("x: ", x, "y: ", y)

	z := append(x, 4)

	fmt.Println(y, z) // ??????
}
