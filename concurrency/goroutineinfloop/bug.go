package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	println(threads)

	wg := sync.WaitGroup{}
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()
			for {
				x++
			}
		}()
	}
	wg.Wait()
	fmt.Println(x)
}
