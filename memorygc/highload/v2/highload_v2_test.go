package main

import (
	"runtime/debug"
	"strconv"
	"testing"
)

func benchmarkBufferAppend(b *testing.B, n int) {
	debug.SetGCPercent(-1) // отключаем GC для бенчмарка

	for i := 0; i < b.N; i++ {
		buf := make([]byte, 0, 20) // Предвыделяем достаточно большой буфер

		for j := 0; j < n; j++ {
			// Очищаем буфер
			buf = buf[:0]

			// Добавляем префикс
			buf = append(buf, 'i')

			// Преобразуем число в строку и добавляем его в буфер
			buf = strconv.AppendInt(buf, int64(j), 10)

			// Обычно в бенчмарках вывод в консоль не используется,
			// так как он может существенно повлиять на результаты измерений.
		}
	}
}

// Определение самого бенчмарка
func BenchmarkBufferAppend10000000(b *testing.B) {
	benchmarkBufferAppend(b, 10000000)
}
