package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

// Профайлер публикует следующие endpoints (см внутри pprof пакет)
// http.HandleFunc("/debug/pprof/", Index)
// http.HandleFunc("/debug/pprof/cmdline", Cmdline)
// http.HandleFunc("/debug/pprof/profile", Profile)
// http.HandleFunc("/debug/pprof/symbol", Symbol)
// http.HandleFunc("/debug/pprof/trace", Trace)
func main() {
	group := sync.WaitGroup{}
	group.Add(1)
	go func() { // http://localhost:6060/debug/pprof/
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			log.Print("error: ", err.Error())
		}
	}()

	// your application code
	group.Wait()
}
