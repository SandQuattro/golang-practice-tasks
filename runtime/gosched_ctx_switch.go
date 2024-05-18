package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 100; i++ {
		runtime.Gosched()
		fmt.Print(s)
	}
}

func main() {
	// установив в 1 явно видно, как переключается контекст и видна конкурентная работа, не параллельная
	// для параллелизма надо установить в > 1
	runtime.GOMAXPROCS(1)
	go say("|")
	go say("*")
	time.Sleep(1 * time.Second)
}
