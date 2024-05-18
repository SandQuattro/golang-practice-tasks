package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

// Функция для обработки запросов
func handleRequest(sem *semaphore.Weighted, requestWeight int64, requestNumber int) {
	// Запрашиваем разрешение на выполнение запроса
	err := sem.Acquire(context.Background(), requestWeight)
	if err != nil {
		fmt.Printf("Запрос №%d не может быть выполнен: %v\n", requestNumber, err)
		return
	}

	// Симулируем обработку запроса
	fmt.Printf("Запрос №%d начал выполнение с весом %d\n", requestNumber, requestWeight)
	time.Sleep(2 * time.Second) // Имитируем время обработки запроса
	fmt.Printf("Запрос №%d завершил выполнение\n", requestNumber)

	// Освобождаем семафор
	sem.Release(requestWeight)
}

func main() {
	// Создаем взвешенный семафор с максимальным весом, равным 10
	sem := semaphore.NewWeighted(10)

	// Запускаем несколько запросов с различными весами
	for i := 0; i < 5; i++ {
		go handleRequest(sem, 1, i) // Обычные запросы
	}

	// Тяжелый запрос
	go handleRequest(sem, 5, 5)

	// Подождем, пока все запросы не будут обработаны
	time.Sleep(10 * time.Second)
}
