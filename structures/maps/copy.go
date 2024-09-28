package main

import "fmt"

func main() {
	m1 := map[string]int{"test": 1}
	fmt.Printf("m1 pointer: %p, m1 value: %v\n", m1, m1)
	m2 := copyMap2(m1)
	fmt.Printf("m1 pointer: %p, m1 value: %v\n", m2, m2)
	fmt.Printf("m2 pointer: %p, m2 value: %v\n", m2, m2)
}

func copyMap(m map[string]int) map[string]int {
	res := make(map[string]int)
	for k, v := range m {
		res[k] = v
	}
	return res
}

func copyMap2(m map[string]int) map[string]int {
	m["test2"] = 2
	return m
}
