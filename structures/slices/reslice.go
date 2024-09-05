package main

import "fmt"

func main() {
	sl := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(sl)
	fmt.Println("sl len:", len(sl), " cap:", cap(sl))

	sl2 := sl[2:4]
	fmt.Println(sl2)
	fmt.Println("sl len:", len(sl2), " cap:", cap(sl2))

	sl3 := sl2[:1]
	fmt.Println(sl3)
	fmt.Println("sl len:", len(sl3), " cap:", cap(sl3))
}
