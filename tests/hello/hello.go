package main

import "fmt"

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return "Hello, world"
}

func localHello() string {
	return "Hello, local world"
}
