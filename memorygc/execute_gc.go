package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

type MyStruct struct {
	data [1024 * 1024]byte // 1 MB данных
}

func main() {
	// отслеживаем изменение в куче

	// Запись в trace файл
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	for i := 0; i < 10; i++ {
		_ = new(MyStruct)
		fmt.Printf("Allocated %d MB\n", (i + 1))
	}

	// Вызов сборщика мусора вручную
	runtime.GC()
	fmt.Println("Garbage collector invoked")

	time.Sleep(2 * time.Second) // Даем время для работы сборщика мусора
}
