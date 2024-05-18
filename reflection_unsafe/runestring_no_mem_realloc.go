package main

import (
	"fmt"
	"unsafe"
)

func main() {
	bytes := []rune{'П', 'р', 'и', 'в', 'е', 'т'}
	fmt.Printf("source bytes: %v\n", bytes)

	strFromBytes := *(*string)(unsafe.Pointer(&bytes))
	fmt.Printf("string from []byte without allocations: %s\n", strFromBytes)

	fmt.Printf("now we reversing string...\n")
	bytes[5] = 'П'
	bytes[4] = 'р'
	bytes[3] = 'и'
	bytes[2] = 'в'
	bytes[1] = 'е'
	bytes[0] = 'т'

	fmt.Printf("source string: %s\n", s)
	fmt.Printf("string from []byte without allocations: %s\n", strFromBytes)
	// fedcba
}
