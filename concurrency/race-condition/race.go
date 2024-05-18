package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Race Conditions
// A race condition occurs when two or more operations must execute in the correct order, but the program
// has not been written so that this order is guaranteed to be maintained.
// Most of the time, this shows up in what’s called a data race, where one concurrent operation attempts to read a
// variable while at some undetermined time another con‐ current operation is attempting to write to the same variable.

// Here, lines 20 and 22 are both trying to access the variable data, but there is no guarantee what order this might
// happen in. There are three possible outcomes to running this code:
// • Nothing is printed. In this case, line 3 was executed before line 5.
// • “the value is 0” is printed. In this case, lines 5 and 6 were executed before line 3.
// • “the value is 1” is printed. In this case, line 5 was executed before line 3, but line 3 was executed before line 6.

func main() {
	goodCodeMutex()
}

// Result can be
// the value is 0.
// the value is 1.
// or
// the value is 0.
// the value is 0.

func badCode() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
	fmt.Printf("the value is %v.\n", data)
}

func goodCodeAtomic() {
	var data int64
	go func() {
		atomic.AddInt64(&data, 1)
	}()
	if atomic.LoadInt64(&data) == 0 {
		fmt.Printf("the value is %v.\n", atomic.LoadInt64(&data))
	}
	fmt.Printf("the value is %v.\n", atomic.LoadInt64(&data))
}

func goodCodeMutex() {
	var data int64
	var mu sync.Mutex

	go func() {
		mu.Lock()
		defer mu.Unlock()
		data++
	}()
	mu.Lock()
	fmt.Printf("the value is %v.\n", data)
	mu.Unlock()
	time.Sleep(5 * time.Millisecond)
	mu.Lock()
	fmt.Printf("the value is %v.\n", data)
	mu.Unlock()
}

func goodCodeWaitGroupsMutex() {
	var ma sync.Mutex
	var value int
	wg := sync.WaitGroup{}
	wg.Add(1)

	ma.Lock()

	go func() {
		ma.Lock()
		defer ma.Unlock()
		value++
		wg.Done()
	}()

	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	}
	ma.Unlock()
	wg.Wait()

	fmt.Printf("the value is %v.\n", value)
}
