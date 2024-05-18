package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	i, j := 42, 2701

	p := &i         // point to i (memory address)
	fmt.Println(p)  // show pointer
	fmt.Println(*p) // read i through the pointer

	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p /= 37       // divide j through the pointer
	fmt.Println(j) // see the new value of j

	var abc string
	var bca string = "Тестовая строка"

	fmt.Printf("size of string abc %d\n", unsafe.Sizeof(abc))
	fmt.Printf("size of string bca %d\n", unsafe.Sizeof(bca))

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&abc))
	fmt.Printf("size of header abc %d\n", unsafe.Sizeof(hdr))

	pointer := *(*string)(unsafe.Pointer(&hdr.Data))
	fmt.Printf("pointer to header abc data %v\n", pointer)

	hdr = (*reflect.StringHeader)(unsafe.Pointer(&bca))
	fmt.Printf("size of header bca %d\n", unsafe.Sizeof(hdr))

	pointer = *(*string)(unsafe.Pointer(&hdr.Data))
	fmt.Printf("pointer to header bca data %v\n", pointer)

	fmt.Printf("pointer to the underlying bytes of str bca %v\n", string(*unsafe.StringData(bca)))
}
