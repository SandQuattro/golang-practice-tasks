package main

import "fmt"

type Parent struct{}

type Child struct {
	Parent
}

func (p *Parent) Print() {
	fmt.Println("i am parent")
}

func main() {
	p := Parent{}
	p.Print() // выведет "i am parent"
	c := Child{}
	c.Print() // выведет "i am parent"
}
