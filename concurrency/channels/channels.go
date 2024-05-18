package main

import (
	"fmt"
)

func main() {
	simpleChannel()
	closeChannelReaderSide()
}

func simpleChannel() {
	ch := make(chan int)
	go func() {
		// создаем горутину и пишем в ней в канал
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	// читаем из канала уже в другой горутине - main
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println(value) // Output: 42
	}
	fmt.Println("simple channel done")
}

func closeChannelReaderSide() {
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		defer close(ch) // закрываем канал done после выхода из функции
		for val := range ch {
			fmt.Println(val)
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
	}

	<-done // ждем, пока читатель прочитает все данные и закроет done
	close(ch)
	fmt.Println("Канал закрыт")
}
