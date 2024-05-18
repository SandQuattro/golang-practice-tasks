package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RateLimiter struct {
	capacity      int64
	tokens        int64
	rate          time.Duration
	lastTokenTime time.Time
	mu            sync.Mutex
}

func NewRateLimiter(capacity int64, rate time.Duration) *RateLimiter {
	return &RateLimiter{
		capacity:      capacity,
		tokens:        capacity,
		rate:          rate,
		lastTokenTime: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastTokenTime)

	// Восполнение токенов
	tokensToAdd := int64(elapsed / rl.rate)
	if tokensToAdd > 0 {
		rl.tokens = min(rl.capacity, rl.tokens+tokensToAdd)
		rl.lastTokenTime = now
	}

	// Проверка доступности токенов
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	limiter := NewRateLimiter(5, time.Second)

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		i := i
		wg.Add(1)
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		go func() {
			defer wg.Done()
			fmt.Printf("Request %d: %t\n", i+1, limiter.Allow())
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()
}
