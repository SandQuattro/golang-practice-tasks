package main

import "fmt"

func main() {
	slice := make([]int, 17, 17)
	slice = append(slice, 0)
	fmt.Printf("len=%d, cap=%d", len(slice), cap(slice)) // 18, 36
}
