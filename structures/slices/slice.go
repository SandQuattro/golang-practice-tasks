package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s []string
	fmt.Printf("s:%v len:%d cap:%d str:%v\n", s, len(s), cap(s), unsafe.SliceData(s))

	s = append(s, "a")
	fmt.Printf("s:%v len:%d cap:%d, str:%v\n", s, len(s), cap(s), unsafe.SliceData(s))
	fmt.Printf("s pointer after append:%p\n", &s)

	strings := make([]string, 5)
	fmt.Printf("strings zero:%s\n", strings[0])

	ptr := &strings[0]
	fmt.Printf("strings zero:%v\n", ptr)

	adddata := []string{"1", "2", "3", "4", "5"}
	fmt.Printf("strings zero:%s\n", adddata[4])

	newstrings := append(strings, adddata...)
	fmt.Printf("Новый слайс %s\n", newstrings)
	//Новый слайс [     1 2 3 4 5]

	strings = append([]string(nil), "")
	fmt.Printf(strings[0])

	strings = append([]string(nil), []string(nil)...)

	for i := range append([]string(nil), []string(nil)...) {
		fmt.Printf("res: %d", i)
	}

	sPtr := make([]string, 1)
	sPtr[0] = "a"
	fmt.Printf("sl pointer:%p\n", &sPtr)
	fmt.Println(sPtr)
	changeSlice(sPtr)
	fmt.Println(sPtr)

}

func changeSlice(sl []string) {
	// здесь уже структура была скопирована, указатели разные, но пока массивы одинаковые
	// до выполнения функции append
	fmt.Printf("sl pointer:%p\n", &sl)
	sl[0] = "b"
	sl = append(sl, "c")
	fmt.Printf("sl pointer:%p\n", &sl)
}
