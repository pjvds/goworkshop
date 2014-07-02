package main

import "fmt"

func start(c chan int) {
	c <- 1
}

func whisper(from, to chan int) {
	to <- 1 + <-from
}

func main() {
	from := make(chan int)
	to := make(chan int)

	go whisper(from, to)

	go start(from)

	fmt.Println(<-to)
}
