package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		fmt.Println("goroutine1")
		return errors.New("err1")
	})

	g.Go(func() error {
		fmt.Println("goroutine2")
		return errors.New("err2")
	})

	// Wait blocks until all function calls from the Go method have returned, then
	// returns the first non-nil error (if any) from them.
	err := g.Wait()
	if err != nil {
		fmt.Println("goroutine error, ", err)
	}
}
