package main

import (
	"fmt"
	"sync"
	"time"
)

// Starvation is any situation where a concurrent process cannot get all the resources it needs to perform work.
// When we discussed livelocks, the resource each goroutine was starved of was a shared lock.
// Livelocks warrant discussion separate from starvation because in a live‐ lock, all the concurrent processes are
// starved equally, and no work is accomplished. More broadly, starvation usually implies that there are one or more
// greedy concur‐ rent process that are unfairly preventing one or more concurrent processes from accomplishing work
// as efficiently as possible, or maybe at all.
func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second
	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
