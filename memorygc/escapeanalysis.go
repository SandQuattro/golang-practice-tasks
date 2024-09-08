package main

// go run -gcflags=-m main.go
// go build -gcflags=-m main.go

// memorygc %  go run -gcflags=-m escapeanalysis.go
// # command-line-arguments
// ./escapeanalysis.go:4:6: can inline main
// ./escapeanalysis.go:8:6: moved to heap: arrayAfter10Mb
// ./escapeanalysis.go:11:23: make([]int, 8192) does not escape
// ./escapeanalysis.go:12:21: make([]int, 8193) escapes to heap
func main() {
	test1()
}

func test1() {
	var arrayBefore10Mb [1310720]int
	arrayBefore10Mb[0] = 1

	var arrayAfter10Mb [1310721]int
	arrayAfter10Mb[0] = 1

	sliceBefore64 := make([]int, 8192)
	sliceOver64 := make([]int, 8193)
	sliceOver64[0] = sliceBefore64[0]

	// Мы видим, что массив `arrayAfter10Mb` был перенесен в кучу, так как его размер превышает 10 МБ,
	// в то время как `arrayBefore10Mb` остался в стеке (для int переменной 10 МБ это 10 * 1024 * 1024 / 8 = 1310720 элементов).
	// Также срез `sliceBefore64` не был отправлен в кучу, поскольку его размер меньше 64 КБ,
	// в то время как `sliceOver64` был сохранен в куче (для int переменной 64 КБ это 64 * 1024 / 8 = 8192 элементов).
	// В общем случае правило такое, если что-то используется в рамках текущей функции, передается как копия,
	// а не по ссылке или из функции возвращается не как ссылка, то оно может быть расположено в стеке
	// Если что-то используется еще и вовне текущей функции, а стало быть возвращается по ссылке, то оно будет перемещено в heap.
}
