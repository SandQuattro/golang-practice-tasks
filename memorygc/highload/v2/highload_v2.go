package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"strconv"
	"sync"
)

var n2 = 10000000

// при N = 10000000 потребление 2.7 мегабайта
func main() {
	// отключаем GC
	debug.SetGCPercent(-1)

	buf := make([]byte, 0, 20) // Предвыделяем достаточно большой буфер

	for i := 0; i < n2; i++ {
		// Очищаем буфер на каждой итерации
		buf = buf[:0]

		// Добавляем префикс
		buf = append(buf, 'i')

		// Преобразуем число в строку и добавляем его в буфер
		buf = strconv.AppendInt(buf, int64(i), 10)

		// Для вывода преобразуем буфер обратно в строку
		if i%100000 == 0 {
			fmt.Println(string(buf))
		}
	}

	group := sync.WaitGroup{}
	group.Add(1)
	go func() { // http://localhost:6060/debug/pprof/
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			log.Print("error: ", err.Error())
		}
	}()
	// your application code
	group.Wait()

}
