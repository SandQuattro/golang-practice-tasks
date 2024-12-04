package main

import "fmt"

func isMatch(string string, pattern string) bool {
	if len(pattern) != len(string) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if string[i] != pattern[i] && pattern[i] != '?' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isMatch("", ""))               // true
	fmt.Println(isMatch("1", "1"))             // true
	fmt.Println(isMatch("1010111", "101?11?")) //true

	fmt.Println(isMatch("", "101?11?"))        // false
	fmt.Println(isMatch("0010101", "101?11?")) // false
}
