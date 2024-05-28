package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var cpu = runtime.NumCPU()

type Counter struct {
	value     atomic.Int32
	alignment [60]byte
}

func (c *Counter) Increment(i int32) {
	c.value.Add(1)
}

func (c *Counter) Get() int32 {
	return c.value.Load()
}

type ShardedCounter struct {
	shards [14]Counter
}

func (c *ShardedCounter) ShardedIncrement(idx int32) {
	c.shards[idx].value.Add(1)
}

func (c *ShardedCounter) ShardedGet(idx int) int32 {
	var value int32
	for i := 0; i < cpu; i++ {
		value += c.shards[idx].Get()
	}
	return value
}

func BenchmarkAtomicCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := Counter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		i := i
		go func() {
			defer wg.Done()
			// bench
			for j := 0; j < b.N; j++ {
				counter.Increment(int32(i))
			}
		}()
	}
	wg.Wait()
}

func BenchmarkShardedCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := ShardedCounter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		i := i
		go func() {
			defer wg.Done()
			// bench
			for j := 0; j < b.N; j++ {
				counter.ShardedIncrement(int32(i))
			}
		}()
	}
	wg.Wait()
}
