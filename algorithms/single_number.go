package main

func singleNumber(arr []int) int {
	singleNum := 0
	for _, num := range arr {
		singleNum ^= num
	}
	return singleNum
}

func main() {
	println(singleNumber([]int{2, 2, 1}) == 1)
	println(singleNumber([]int{4, 1, 2, 1, 2}) == 4)
	println(singleNumber([]int{1}) == 1)
}
