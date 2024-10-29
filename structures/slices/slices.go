package main

import (
	"fmt"
)

func main() {
	a1 := make([]int, 0, 5) // len=0 cap=5
	//a1 := make([]int, 5)  // len=5 cap=5 !!
	a1 = append(a1, []int{1, 2, 3, 4, 5}...)
	a2 := append(a1, 6)
	a3 := append(a2, 7)

	fmt.Println(a1, a2, a3)
	fmt.Printf("a1=%p, a2=%p, a3=%p\n", a1, a2, a3)

	sl := []string{"a", "b", "c"}
	fmt.Println(len(sl), cap(sl))

	sl2 := sl[1:2]
	fmt.Println(len(sl2), cap(sl2))

	add(sl[1:2])
	fmt.Println(sl)

	arr := [3]int{1, 2, 3}
	sl3 := arr[:]
	sl3[0] = 3
	fmt.Println("-------")
	fmt.Println(arr) // <- ?

	arr2 := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	sl4 := arr2[1:]
	sl4[0][0] = 3
	fmt.Println("-------")
	fmt.Println(arr2) // <- ?
	fmt.Println(sl4)  // <- ?

	arr3 := [3]int{1, 2, 3}
	sl5 := arr3[1:2]
	sl5 = append(sl5, 4)
	sl5[0] = 8
	fmt.Println("-------")
	fmt.Println(arr3) // <- ?
	fmt.Println(sl5)  // <- ?

}

func add(s []string) {
	s = append(s, "x")
}
