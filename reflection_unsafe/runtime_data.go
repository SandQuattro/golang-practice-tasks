package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []byte{3, 5, 7}

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s)) // case 1
	fmt.Printf("hdr data: 0x%p\n", &hdr.Data)

	array := *(*[]byte)(unsafe.Pointer(&hdr.Data))
	fmt.Printf("array: %v array addr: %p\n", array, &array)

	sizeof := unsafe.Sizeof(array)
	println(sizeof)

	slice2 := []int{1, 2, 3, 4, 5}
	// de-slicing
	array2 := (*[5]int)(slice2)

	hdr2 := (*reflect.SliceHeader)(unsafe.Pointer(&slice2))

	// указатель на первый элемент, SliceData returns &slice2[:1][0].
	data := unsafe.SliceData(slice2)

	fmt.Printf("slice data %v, SliceData address: 0x%x, address of de-slice Array %p \n", slice2, data, array2)
	fmt.Printf("slice header struct: 0x%x len:%d capacity:%d \n", hdr2.Data, hdr2.Len, hdr2.Cap)

	array2[1] = 9

	fmt.Printf("new slice data %v, array data: %v\n", slice2, array2)

	m := make(map[string]any)
	m["tt"] = struct{}{}
	fmt.Printf("map reflect: %v\n", reflect.ValueOf(m))

}
