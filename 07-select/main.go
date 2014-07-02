package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)

	go talk("rob", c)
	go talk("christian", c)

	for {
		v := <-c
		fmt.Printf(v)
	}
}

func talk(name string, c chan string) {
	for {
		c <- fmt.Sprintf("my name is %v\n", name)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}
