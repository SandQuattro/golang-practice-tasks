package main

import "fmt"

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	fmt.Println(fmt.Sprintf("s=%v,len=%d, cap=%d", s, len(s), cap(s)))

	// Делим слайс пополам:
	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]
	fmt.Println(fmt.Sprintf("s1=%v,len=%d, cap=%d", s1, len(s1), cap(s1)))
	fmt.Println(fmt.Sprintf("s2=%v,len=%d, cap=%d", s2, len(s2), cap(s2)))

	fmt.Println("модифицируем элементы слайсов")

	s1[0] = 32
	s1[1] = 33
	s1[2] = 34

	s2[0] = 42
	s2[1] = 43
	s2[2] = 44

	fmt.Println(fmt.Sprintf("s1=%v,len=%d, cap=%d", s1, len(s1), cap(s1)))
	fmt.Println(fmt.Sprintf("s2=%v,len=%d, cap=%d", s2, len(s2), cap(s2)))

	fmt.Println(fmt.Sprintf("s=%v,len=%d, cap=%d", s, len(s), cap(s)))

}
