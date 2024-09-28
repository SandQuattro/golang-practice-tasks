package main

import "fmt"

func main() {
	m1 := map[string]int{"test": 1}
	fmt.Printf("m2 pointer: %p, value: %v\n", m1, m1)
	m2 := copyMap(m1)
	fmt.Printf("m2 pointer: %p, value: %v\n", m2, m2)
}

func copyMap(m map[string]int) map[string]int {
	res := make(map[string]int)
	for k, v := range m {
		res[k] = v
	}
	return res
}
