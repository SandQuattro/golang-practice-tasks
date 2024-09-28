package main

import (
	"fmt"
	"time"
)

// WaitMany waits for a and b to close.
func WaitMany(a, b chan bool) {
	var aclosed, bclosed bool
	for !aclosed || !bclosed {
		select {
		case <-a:
			fmt.Println("A")
			aclosed = true
		case <-b:
			fmt.Println("B")
			bclosed = true
		}
	}
}

func WaitManyFixed(a, b chan bool) {
	for a != nil || b != nil {
		select {
		case <-a:
			fmt.Println("A")
			a = nil
		case <-b:
			fmt.Println("B")
			b = nil
		}
	}
}

func main() {
	a, b := make(chan bool), make(chan bool)
	t0 := time.Now()

	go func() {
		close(a)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		close(b)
	}()

	// WaitMany(a, b)
	WaitManyFixed(a, b)
	fmt.Printf("Waited %v for WaitMany\n", time.Since(t0))
}
