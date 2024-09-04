package main

func recursiveFunction(sl []int) {
	recursiveFunction(append(sl, 1))
}

func main() {
	sl := make([]int, 0)
	recursiveFunction(sl)
}
