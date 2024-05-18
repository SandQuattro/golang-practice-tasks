package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	user := struct {
		balance       float64
		isTestProfile bool
		age           uint64
		isDesktop     bool
	}{}
	fmt.Println(unsafe.Sizeof(user.balance))       //
	fmt.Println(unsafe.Sizeof(user.age))           //
	fmt.Println(unsafe.Sizeof(user.isTestProfile)) //
	fmt.Println(unsafe.Sizeof(user.isDesktop))     //
	fmt.Println(unsafe.Sizeof(user))               //
	fmt.Println("----")

	i := new(int)
	*i = 10

	ptr := unsafe.Pointer(&user)
	fmt.Println(unsafe.Sizeof(&ptr)) //

	fmt.Println(unsafe.Sizeof([]struct{}{})) //
	fmt.Println(unsafe.Sizeof([5]int64{}))   //
	fmt.Println(unsafe.Sizeof("h"))          //

	fmt.Println("----")
	var longStr string
	for i := 0; i <= 100; i++ {
		longStr += strconv.Itoa(i)
	}
	fmt.Println(unsafe.Sizeof(longStr)) //

	fmt.Println("----")
	fmt.Println(unsafe.Sizeof(make(map[string][]string))) //
}
