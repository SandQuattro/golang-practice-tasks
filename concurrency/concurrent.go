package main

import (
	"fmt"
	"runtime"
	"sync"
)

// почему здесь вывод 4 0 1 2 3? ответ - см шедулер, local run queue, fifo + 1-element lifo stack
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i) // Передаем i как параметр
	}

	wg.Wait()
}
