package main

import "fmt"

// Sliding window
func lengthWithoutRepeats(s string) int {
	answer, left, right := 0, 0, 0
	uniqueChars := make(map[rune]bool)

	// do until right pointer did not get to finish
	for right < len(s) {
		// here we are expanding window
		if !uniqueChars[rune(s[right])] {
			// add char to map
			uniqueChars[rune(s[right])] = true
			// updating our result to actual length
			answer = max(answer, right-left+1)
			right++
		} else {
			// here we collapse our window
			delete(uniqueChars, rune(s[left]))
			left++
		}
	}

	return answer
}

func main() {
	fmt.Println(lengthWithoutRepeats("abcbada"))   // 4
	fmt.Println(lengthWithoutRepeats("abcbdbefg")) // 5
	fmt.Println(lengthWithoutRepeats("axbxcxd"))   // 3
	fmt.Println(lengthWithoutRepeats("aaaaaaa"))   // 1
	fmt.Println(lengthWithoutRepeats("abcdefg"))   // 7

}
