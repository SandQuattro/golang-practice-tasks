package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
)

type MyStruct struct {
	data [1024 * 1024]byte // 1 MB данных
}

// go tool trace trace.out
func main() {
	// отслеживаем изменение в куче
	debug.SetGCPercent(-1)
	//debug.SetGCPercent(10)
	//debug.SetGCPercent(100)
	//debug.SetGCPercent(1000)

	// Запись в trace файл
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	for i := 0; i < 100; i++ {
		s := new(MyStruct)
		for j := 0; j < len(s.data); j++ {
			s.data[j] = byte(i)
		}
		fmt.Printf("Allocated %d MB\n", i+1)

		if i == 50 {
			// Вызов сборщика мусора вручную
			runtime.GC()
			fmt.Println("Garbage collector invoked")
		}
	}

}
