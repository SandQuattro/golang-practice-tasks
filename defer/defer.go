package main

import "fmt"

// defer - выполнение функции после выхода из функции
// фактически выполняется прямо перед выходом из функции, до возврата значения функцией return
func main() {
	fmt.Println(test1())
	test2()
}

func test1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func test2() {
	var i1 = 10
	var v2 = 5
	var i2 = &v2

	// Значения рассчитываются на момент объявления defer
	defer fmt.Println("i2=", *i2)
	defer fmt.Println("i1=", i1)
	i1 = 20
	v2 = 100
}
