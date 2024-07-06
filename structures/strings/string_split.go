package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "1.1.1.1/24"
	split := strings.Split(str, ".")
	fmt.Println(split)
	for _, str := range split {
		sp := strings.Split(str, "/")
		fmt.Println(sp)
	}
}
