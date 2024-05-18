package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		mutex.Lock()
		go func(ii int) {
			defer mutex.Unlock()
			defer wg.Done()
			c <- fmt.Sprintf("%d", ii)
		}(i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for b := range c {
		println(b)
	}

}
