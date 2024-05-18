package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(startWith("приве", "п"))
	fmt.Println(startWith2("приве", "п"))
	fmt.Println(startWith3("😄приве", "😄"))
	fmt.Println(startWith4("😄приве", "😄"))
}

func startWith(in, test string) bool {
	if len(test) > len(in) {
		return false
	}

	i := 0
	for range test {
		if test[i] != in[i] {
			return false
		}
		i++
	}

	return true
}

func startWith2(in, test string) bool {
	return strings.HasPrefix(in, test)
}

func startWith3(in, test string) bool {
	for i, _ := range test {
		if test[i] != in[i] {
			return false
		}
	}
	return true
}

func startWith4(in, test string) bool {
	for i, val := range test {
		r, _ := utf8.DecodeRune([]byte(in[i:]))
		if val != r {
			return false
		}
	}
	return true
}
