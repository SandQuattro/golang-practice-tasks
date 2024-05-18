package main

// запись в nil канал приводит к deadlock, nil каналы не готовы к обмену данными
// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [chan receive (nil chan)]:
// goroutine 17 [chan send (nil chan)]:

func main() {
	var c chan string
	c <- "let's get started" // deadlock
}
