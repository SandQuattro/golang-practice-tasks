package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

type Counter struct {
	value     atomic.Int32
	alignment [60]byte
}

func (c *Counter) Increment() {
	c.value.Add(1)
}

func (c *Counter) Get() int32 {
	return c.value.Load()
}

type ShardedCounter struct {
	shards [10]Counter
}

func (c *ShardedCounter) ShardedIncrement(idx int) {
	c.shards[idx].value.Add(1)
}

func (c *ShardedCounter) ShardedGet(idx int) int32 {
	var value int32
	for i := 0; i < 10; i++ {
		value += c.shards[idx].Get()
	}
	return value
}

func BenchmarkAtomicCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := ShardedCounter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		i := i
		go func() {
			defer wg.Done()
			// bench
			for j := 0; j < b.N; j++ {
				counter.ShardedIncrement(i)
			}
		}()
	}
	wg.Wait()
}
