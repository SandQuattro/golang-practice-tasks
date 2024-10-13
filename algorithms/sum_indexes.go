package main

import "fmt"

// time: O(n), mem: O(n)
func twoSum(nums []int, target int) []int {
	// map[num]index
	mem := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// target = firstTera + secondTerm
		// return firstTeraIndex, secondTermIndex
		firstTerm := nums[i]
		firstTermIndex := i

		secondTerm := target - firstTerm

		secondTermIndex, exists := mem[secondTerm]
		if exists {
			return []int{secondTermIndex, firstTermIndex}
		}
		mem[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 18))
}
