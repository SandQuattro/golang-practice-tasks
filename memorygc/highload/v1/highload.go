package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"sync"
)

// Профайлер публикует следующие endpoints (см внутри pprof пакет)
// http.HandleFunc("/debug/pprof/", Index)
// http.HandleFunc("/debug/pprof/cmdline", Cmdline)
// http.HandleFunc("/debug/pprof/profile", Profile)
// http.HandleFunc("/debug/pprof/symbol", Symbol)
// http.HandleFunc("/debug/pprof/trace", Trace)

var bufferPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

const maxBufferSize = 64 // Максимальный размер буфера в байтах, например, 64 байта

func getBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func putBuffer(buf *bytes.Buffer) {
	if buf.Cap() > maxBufferSize {
		// Если размер буфера больше установленного предела, то не возвращаем его в пул
		return
	}
	buf.Reset() // Очищаем содержимое буфера перед возвращением в пул
	bufferPool.Put(buf)
}

const n1 = 10000000

// при N = 10000000 потребление 165 мегабайт
func main() {
	// отключаем GC
	debug.SetGCPercent(-1)

	go func() {
		for i := 0; i < n1; i++ {
			// берем кусок из пула
			buf := getBuffer()
			buf.WriteString(fmt.Sprintf("i%d", i))
			if i%100000 == 0 {
				fmt.Println(buf.String())
			}
			// кладем в пул отработанный кусок
			putBuffer(buf)
		}
	}()

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
