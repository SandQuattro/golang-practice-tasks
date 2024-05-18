package main

import (
	"fmt"
)

type CircularQueue struct {
	items            [10]int
	head, tail, size int
}

func NewCircularQueue() *CircularQueue {
	return &CircularQueue{}
}

func (q *CircularQueue) Push(item int) {
	if q.size == len(q.items) {
		// Если очередь полная, переписываем самый старый элемент
		q.tail = (q.tail + 1) % len(q.items)
	} else {
		q.size++
	}
	q.items[q.head] = item
	q.head = (q.head + 1) % len(q.items)
}

func (q *CircularQueue) Print() {
	for i := 0; i < q.size; i++ {
		idx := (q.tail + i) % len(q.items)
		fmt.Printf("%d ", q.items[idx])
	}
	fmt.Println()
}

func main() {
	queue := NewCircularQueue()
	for i := 1; i <= 50; i++ {
		queue.Push(i)
		queue.Print() // Вывод текущего состояния очереди
	}
}
