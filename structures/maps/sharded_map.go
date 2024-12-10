package main

import (
	"fmt"
	"sync"
)

type ShardedMap struct {
	shards    []*Shard
	numShards int
}

type Shard struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewShardedMap(numShards int) *ShardedMap {
	sm := &ShardedMap{
		shards:    make([]*Shard, numShards),
		numShards: numShards,
	}

	for i := 0; i < numShards; i++ {
		sm.shards[i] = &Shard{
			data: make(map[string]interface{}),
		}
	}

	return sm
}

func (sm *ShardedMap) getShard(key string) *Shard {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sm.shards[sum%sm.numShards]
}

func (sm *ShardedMap) Set(key string, value interface{}) {
	shard := sm.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = value
}

func (sm *ShardedMap) Get(key string) (interface{}, bool) {
	shard := sm.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	val, ok := shard.data[key]
	return val, ok
}

func (sm *ShardedMap) Delete(key string) {
	shard := sm.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	delete(shard.data, key)
}

func main() {
	sm := NewShardedMap(4)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			sm.Set(key, i)
		}(i)

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			if val, ok := sm.Get(key); ok {
				fmt.Printf("Key: %s, Value: %v\n", key, val)
			}
		}(i)
	}

	//for i := 0; i < 100; i++ {
	//
	//}

	wg.Wait()
}
