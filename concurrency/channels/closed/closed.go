package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	// после закрытия канала читатели возвращают default значение для типа канала (int = 0..),
	// писатели - паникуют
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// panic: send on closed channel
	ch <- 0
}
