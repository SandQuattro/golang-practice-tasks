package main

import (
	"context"
	"time"
)

func main() {

	// сценарии использования контекста
	ctx := context.WithValue(context.Background(), 1, 2)

	ctx2, cancel := context.WithCancel(ctx)
	cancel()

	ctx3, cancel2 := context.WithTimeout(ctx2, 5*time.Second)
	cancel2()

	ctx4, cancel3 := context.WithDeadline(ctx3, time.Now().Add(5*time.Second))
	cancel3()

	//<-ctx.Done() -- deadlock, we newer will wait for this ctx cancel
	<-ctx2.Done()
	<-ctx3.Done()
	<-ctx4.Done()

}
