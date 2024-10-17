package main

import "fmt"

func isMatch(string string, pattern string) bool {
	return true
}

func main() {
	fmt.Println(isMatch("", ""))               // true
	fmt.Println(isMatch("1", "1"))             // true
	fmt.Println(isMatch("1010111", "101?11?")) //true

	fmt.Println(isMatch("", "101?11?"))        // false
	fmt.Println(isMatch("0010101", "101?11?")) // false
}
