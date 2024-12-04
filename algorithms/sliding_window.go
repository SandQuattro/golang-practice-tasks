package main

import "fmt"

// Sliding window
func lengthWithoutRepeats(s string) int {
	left, right := 0, 0
	uniqueChars := make(map[rune]bool)
	maxLength := 0

	for right < len(s) {
		// here ww are expanding window
		if !uniqueChars[rune(s[right])] {
			uniqueChars[rune(s[right])] = true
			right++
			maxLength = max(maxLength, len(uniqueChars))
		} else {
			// here we collapse our window
			delete(uniqueChars, rune(s[left]))
			left++
		}
	}

	return maxLength
}

func main() {
	fmt.Println(lengthWithoutRepeats("abcbdbefg"))
}
