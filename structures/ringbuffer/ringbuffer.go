package main

import (
	"fmt"
)

type RingBuffer struct {
	buffer []int
	size   int
	start  int
	end    int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]int, size),
		size:   size,
		start:  0,
		end:    -1,
	}
}

func (rb *RingBuffer) Push(value int) {
	rb.end = (rb.end + 1) % rb.size
	rb.buffer[rb.end] = value

	if rb.end == rb.start && rb.end != -1 {
		rb.start = (rb.start + 1) % rb.size // Move start forward if we are overwriting
	}
}

func (rb *RingBuffer) Pop() (int, bool) {
	if rb.end == -1 {
		return 0, false // Buffer is empty
	}

	value := rb.buffer[rb.start]
	if rb.start == rb.end {
		rb.start, rb.end = 0, -1 // Reset buffer if it was the last element
	} else {
		rb.start = (rb.start + 1) % rb.size
	}
	return value, true
}

func main() {
	rb := NewRingBuffer(5)

	// Adding elements to the buffer
	for i := 1; i <= 5; i++ {
		rb.Push(i)
		fmt.Println("Pushed:", i)
	}

	// Trying to add more elements than the buffer can hold
	rb.Push(6)
	fmt.Println("Pushed:", 6)

	// Popping elements from the buffer
	for i := 1; i <= 5; i++ {
		if val, ok := rb.Pop(); ok {
			fmt.Println("Popped:", val)
		}
	}
}
