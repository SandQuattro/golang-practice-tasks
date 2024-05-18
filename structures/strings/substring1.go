package main

import (
	"fmt"
)

func main() {
	str := "Hello, 世界"
	str2 := str[0:5]
	//sl[0] = 'X' - нельзя изменить строку
	fmt.Printf("%s\n%s\n %p\n", str, str2, &str2)

	sl := []byte(str2)
	fmt.Printf("slice %s\n %p\n", sl, &sl)
	sl[0] = 'X'
	fmt.Printf("changed slice %s\n %p", sl, &sl)
	fmt.Println("-------")
	fmt.Printf("%s\n", str)
	fmt.Printf("%s\n", str2)
}
