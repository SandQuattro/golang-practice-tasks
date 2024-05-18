package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

const (
	i = 5
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	// de-slicing
	if len(slice) == i {
		array := (*[i]int)(slice)

		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

		// указатель на первый элемент, SliceData returns &slice[:1][0].
		data := unsafe.SliceData(slice)

		fmt.Printf("slice data %v, SliceData address: 0x%x, address of de-slice Array %p \n", slice, data, array)
		fmt.Printf("slice header struct: 0x%x len:%d capacity:%d \n", hdr.Data, hdr.Len, hdr.Cap)

		array[0] = 9

		fmt.Printf("new slice data %v, array data: %v", slice, array)

		// копирование
		arr := [i]int{}
		copy(arr[:], slice[:])
		fmt.Println(arr)
	}
}
