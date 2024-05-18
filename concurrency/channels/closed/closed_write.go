package main

import "fmt"

func main() {
	// закрываем канал со стороны читателя
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case val, ok := <-ch:
				if !ok {
					close(done)
					return
				}
				fmt.Println(val)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

	<-done // ждем, пока читатель прочитает все данные и закроет done
	fmt.Println("Канал закрыт")
}

func writeOnClosedChan() {
	ch := make(chan string)

	close(ch)

	ch <- "Hello" // тут будет паника
}
