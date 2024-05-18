package main

import "fmt"

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	for _, c := range s {
		m[c]++
	}

	for _, c := range t {
		if m[c] == 0 {
			return false
		}
		m[c]--
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
