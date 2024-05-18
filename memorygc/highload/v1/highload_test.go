package main

import (
	"fmt"
	"testing"
)

func BenchmarkBufferOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := getBuffer()
		buf.WriteString(fmt.Sprintf("i%d", i)) // Пример использования
		putBuffer(buf)
	}
}

// Также можно добавить бенчмарк для отдельных операций, если это необходимо
func BenchmarkGetBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getBuffer()
	}
}

func BenchmarkPutBuffer(b *testing.B) {
	buf := getBuffer()
	b.ResetTimer() // Начинаем отсчет времени заново, чтобы не учитывать время на получение буфера
	for i := 0; i < b.N; i++ {
		putBuffer(buf)
	}
}
