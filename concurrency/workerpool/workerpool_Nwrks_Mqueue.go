package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Task представляет собой функцию, которую нужно выполнить
type Task func() error

// Pool управляет пулом воркеров и очередью задач
type Pool struct {
	maxWorkers   int
	maxQueueSize int
	taskQueue    chan Task
	workersWG    sync.WaitGroup
	workerSem    chan struct{} // Семафор для ограничения количества воркеров
}

var (
	ErrQueueFull  = errors.New("task queue is full")
	ErrPoolClosed = errors.New("worker pool is closed")
)

// New создает новый пул с заданным количеством воркеров и размером очереди
func New(maxWorkers, maxQueueSize int) *Pool {
	return &Pool{
		maxWorkers:   maxWorkers,
		maxQueueSize: maxQueueSize,
		taskQueue:    make(chan Task, maxQueueSize),
		workerSem:    make(chan struct{}, maxWorkers), // Создаем семафор с лимитом воркеров
	}
}

// Submit добавляет новую задачу в очередь
func (p *Pool) Submit(task Task) error {
	select {
	case p.taskQueue <- task:
		p.startWorker()
		return nil
	default:
		return ErrQueueFull
	}
}

func (p *Pool) startWorker() {
	select {
	case p.workerSem <- struct{}{}: // Пытаемся захватить семафор
	default:
		return // Если семафор занят, не создаем нового воркера
	}

	p.workersWG.Add(1)
	go p.process()
}

// worker обрабатывает задачи из очереди
func (p *Pool) process() {
	defer func() {
		<-p.workerSem // Освобождаем семафор
		p.workersWG.Done()
	}()

	for {
		select {
		case task, opened := <-p.taskQueue:
			if !opened {
				return
			}
			_ = task() // Выполняем задачу, игнорируя ошибку
		}
	}
}

// Close закрывает пул и ожидает завершения всех задач
func (p *Pool) Close() {
	close(p.taskQueue)
	p.workersWG.Wait()
}

func main() {
	pool := New(5, 10)

	for i := 0; i < 50; i++ {
		taskNum := i

		err := pool.Submit(func() error {
			fmt.Printf("Starting task %d\n", taskNum)
			// Имитируем долгую работу
			time.Sleep(time.Second)
			fmt.Printf("Finished task %d\n", taskNum)
			return nil
		})

		if err != nil {
			fmt.Printf("Failed to submit task %d: %s\n", taskNum, err)
		}

		// Добавляем небольшую задержку между отправкой задач
		time.Sleep(time.Millisecond * 50)
	}

	// Ждем завершения всех задач
	pool.Close()
	fmt.Println("All tasks completed")
}
