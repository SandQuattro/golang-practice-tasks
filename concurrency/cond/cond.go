package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Cond is a rendezvous for goroutines waiting for or announcing the occurence of an event/signal.

// We can also use Broadcast() instead of Signal(). Broadcast() tells all goroutines that something happened.
// Signal() tells the goroutine that is waiting the longest that something happened.
var (
	cond  = sync.NewCond(&sync.Mutex{})
	value int
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Изменяем значение переменной value через 5 секунд
		<-time.After(5 * time.Second)
		value = 42

		// Сигнализируем всем ожидающим горутинам, что значение изменилось
		cond.Broadcast()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Print("goRoutine1 started")
		// Ждем, пока значение переменной value изменится
		cond.L.Lock()
		if value == 0 {
			cond.Wait()
		}
		cond.L.Unlock()

		// Выводим значение переменной value
		fmt.Println("goRoutine1, Value is", value)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Print("goRoutine2 started")
		// Ждем, пока значение переменной value изменится
		cond.L.Lock()
		if value == 0 {
			cond.Wait()
		}
		cond.L.Unlock()

		// Выводим значение переменной value
		fmt.Println("goRoutine2", " Value is", value)
	}()
	wg.Wait()
}
