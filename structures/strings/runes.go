package main

import "fmt"

func main() {
	str := "asdf🍀"
	for _, runeValue := range str {
		fmt.Printf("%U\n", runeValue)
	}
}
