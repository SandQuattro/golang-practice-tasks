package main

import (
	"fmt"
	"unsafe"
)

func main() {
	bytes := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	fmt.Printf("source string: %s\n", bytes)

	strFromBytes := *(*string)(unsafe.Pointer(&bytes))
	fmt.Printf("string from []byte without allocations: %s\n", strFromBytes)

	fmt.Printf("now we reversing string...\n")
	bytes[5] = 'a'
	bytes[4] = 'b'
	bytes[3] = 'c'
	bytes[2] = 'd'
	bytes[1] = 'e'
	bytes[0] = 'f'
	fmt.Printf("string from []byte without allocations: %s\n", strFromBytes)
	// fedcba

	// Но если в вышеуказанный слайс байт присвоить новый слайс,
	// то будет переаллокация памяти, и наша строка и новый слайс байт смотрят уже в разные.
	// Пример:
	bytes = []byte{'a', 'a', 'a', 'a', 'a', 'a'}
	fmt.Printf("string from []byte without allocations: %s\n", strFromBytes)
	// fedcba
}
