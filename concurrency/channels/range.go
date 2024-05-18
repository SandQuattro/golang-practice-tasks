package main

func main() {
	ch := make(chan int)
	go func() {
		close(ch)
	}()
	for val := range ch {
		println(val)
	}
}
