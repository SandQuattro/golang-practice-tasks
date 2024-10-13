package main

import (
	"fmt"
)

func sumArrays(arr1, arr2 []int) []int {
	resultLen := max(len(arr1), len(arr2)) + 1
	result := make([]int, resultLen)

	carry := 0
	for i := 0; i < resultLen; i++ {
		sum := 0
		sum += carry

		if i < len(arr1) {
			sum += arr1[len(arr1)-1-i]
		}

		if i < len(arr2) {
			sum += arr2[len(arr2)-1-i]
		}

		result[resultLen-1-i] = sum % 10
		carry = sum / 10
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(sumArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(sumArrays([]int{5, 7}, []int{9}))
}
