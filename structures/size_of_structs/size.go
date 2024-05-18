package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type t struct {
	err error
}

type Order struct {
	Premium bool    // 0
	NewFlow bool    // 8 тут сработает оптимизация, 2 байта будет занято обоими bool, остальные 6 байт - 0
	UserID  int64   // 8
	ItemID  int64   // 8
	Price   float64 // 8
	//Status  string  // 16 байт( unsafe.Sizeof(reflect.StringHeader{}) )
}

func main() {
	var abc string
	var bca string = "Тестовая строка"

	fmt.Printf("size of string abc %d\n", unsafe.Sizeof(abc))
	fmt.Printf("size of string bca %d\n", unsafe.Sizeof(bca))

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&abc))
	fmt.Printf("size of header abc %d\n", unsafe.Sizeof(hdr))

	hdr = (*reflect.StringHeader)(unsafe.Pointer(&bca))
	fmt.Printf("size of header bca %d\n", unsafe.Sizeof(hdr))

	pointer := *(*string)(unsafe.Pointer(&hdr.Data))
	fmt.Printf("pointer to header bca data %v\n", pointer)

	fmt.Printf("size of pointer to string abc %d\n", unsafe.Sizeof(&abc))
	fmt.Printf("size of pointer to string bca %d\n", unsafe.Sizeof(&bca))

	fmt.Printf("size of empty struct: %d, can be used as channel ping signal \n", unsafe.Sizeof(struct{}{}))

	fmt.Printf("size of empty error struct: %d\n", unsafe.Sizeof(t{}))

	var a int32 // 4

	s := []int32{1, 2, 3}
	hdr2 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("size of slice array: %d, len:%d\n", unsafe.Sizeof(a), hdr2.Len)

	fmt.Printf("size of Order struct: %d\n", unsafe.Sizeof(Order{}))

	fmt.Printf("size of string: %d\n", unsafe.Sizeof(reflect.StringHeader{}))
}
