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
}
