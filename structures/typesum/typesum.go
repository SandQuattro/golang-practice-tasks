package main

import "fmt"

type Sum interface {
	Add(Sum) Sum
}

type IntSum struct {
	value int64
}

func main() {
	var a Sum
	a = &IntSum{value: 1}
	a = a.Add(&IntSum{value: 2})
	a = a.Add(&IntSum{value: 3})
	fmt.Println(a)
}

func (i *IntSum) Add(a Sum) Sum {
	switch x := a.(type) {
	case *IntSum:
		return &IntSum{
			value: i.value + x.value,
		}

	default:
		return nil
	}
}
