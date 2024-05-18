package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "тут лежит длинная строка"
	fmt.Println("исходная строка: ", str)
	bytes := []byte(str)

	// получаем доступ к нижележащему массиву
	// строка это reflect.StringHeader
	// Data uintptr
	// Len  int

	pointer := unsafe.Pointer(&bytes)
	fmt.Println("pointer to underlying array: ", pointer)

	strPointer := (*[]byte)(pointer)
	fmt.Println("strPointer to underlying array: ", strPointer)
	// pointer to underlying array:     0x14000096230
	// strPointer to underlying array:  0x14000096230

	//  делаем dereference, в итоге у нас получается тот тип, который указали выше, можем указать *[]byte, *string
	strSlice := *strPointer
	fmt.Println("slice: ", strSlice)
	fmt.Println("string slice: ", string(strSlice))

	strSlice[0] = '.'
	strSlice[1] = '.'
	strSlice[2] = '.'

	fmt.Println("string slice after modification: ", string(strSlice))
}
