package main

import "fmt"

func main() {
	fmt.Println(TernaryOneType(true, 1, 2))
	fmt.Println(TernaryAny(false, 1, "збс"))
}

// Используется только для одного типа
func TernaryOneType[T any](cond bool, ifTrue T, ifFalse T) T {
	if cond {
		return ifTrue
	}
	return ifFalse
}

// Любые типы на вход, и на выход
func TernaryAny(cond bool, ifTrue any, ifFalse any) any {
	if cond {
		return ifTrue
	}
	return ifFalse
}
