package main

import (
	"container/list"
	"context"
	"fmt"
	"log"
	"sync"
)

type MyRedis interface {
	Get(ctx context.Context, key string) (value string, err error)
	Set(key, value string)
}

// LRU Cache structure
type LRU struct {
	capacity  int
	cache     map[string]*list.Element
	evictList *list.List
	mx        sync.RWMutex
}

type entry struct {
	key   string
	value string
}

// NewLRU creates a new LRU cache
func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity:  capacity,
		cache:     make(map[string]*list.Element),
		evictList: list.New(),
	}
}

type Redis struct {
	data map[string]string
	mx   sync.RWMutex
	lru  *LRU
}

// NewRedis creates a new Redis-like instance with LRU caching
func NewRedis(capacity int) *Redis {
	return &Redis{
		data: make(map[string]string),
		lru:  NewLRU(capacity),
	}
}

func (r *Redis) Get(ctx context.Context, key string) (value string, err error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	// Try to get from LRU cache first
	if val, ok := r.lru.Get(key); ok {
		return val, nil
	}

	// If not found in LRU, look in main storage
	if val, ok := r.data[key]; ok {
		// Update LRU
		r.lru.Set(key, val)
		return val, nil
	}
	return "", nil
}

func (r *Redis) Set(key, value string) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.data[key] = value
	r.lru.Set(key, value)
}

// Get retrieves a value from the cache
func (l *LRU) Get(key string) (string, bool) {
	l.mx.RLock()
	defer l.mx.RUnlock()

	if ele, ok := l.cache[key]; ok {
		l.evictList.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return "", false
}

// Set adds a value to the cache
func (l *LRU) Set(key, value string) {
	l.mx.Lock()
	defer l.mx.Unlock()

	if ele, ok := l.cache[key]; ok {
		l.evictList.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}

	ele := l.evictList.PushFront(&entry{key, value})
	l.cache[key] = ele

	if l.evictList.Len() > l.capacity {
		l.removeOldest()
	}
}

// removeOldest removes the oldest item from the cache
func (l *LRU) removeOldest() {
	ele := l.evictList.Back()
	if ele != nil {
		l.evictList.Remove(ele)
		delete(l.cache, ele.Value.(*entry).key)
	}
}

func main() {
	redis := NewRedis(2)
	redis.Set("key1", "value1")
	redis.Set("key2", "value2")
	redis.Set("key3", "value3")
	val, err := redis.Get(context.Background(), "key2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}
