package main

import "fmt"

func main() {
	calcFibo(0, 1)
}

func calcFibo(i, j int) {
	if i > 100 {
		return
	}
	fmt.Println(i)
	calcFibo(i+j, i)
}
