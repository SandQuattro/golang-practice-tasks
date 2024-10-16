package main

import "fmt"

type Stack[T comparable] struct {
	data []T
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

func (s *Stack[T]) Pop() {
	if len(s.data) > 0 {
		s.data = s.data[:len(s.data)-1]
	}
}

func main() {
	stack := Stack[int]{
		data: make([]int, 0),
	}

	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack)

	stack.Pop()
	stack.Pop()
	fmt.Println(stack)
}
