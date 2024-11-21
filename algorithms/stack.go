package main

import "fmt"

type MyInt int

// Stack ~ means underlying type int and types based on it
type Stack[T ~int] struct {
	stack []T
}

func (s *Stack[T]) Push(data T) {
	s.stack = append(s.stack, data)
}

func (s *Stack[T]) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *Stack[T]) Top() (T, error) {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1], nil
	}

	return -1, fmt.Errorf("stack is empty")
}

func (s *Stack[T]) IsEmpty() bool {
	if len(s.stack) == 0 {
		return true
	}

	return false
}

func main() {
	// using here underlying type of MyInt
	stack := &Stack[int]{
		stack: make([]int, 0),
	}

	stack.Push(1)

	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.Top())

	// using here MyInt type
	stack2 := &Stack[MyInt]{
		stack: make([]MyInt, 0),
	}
	var a MyInt = 5
	stack2.Push(a)
}
