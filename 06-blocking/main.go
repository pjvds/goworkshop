package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go talk("rob")
	go talk("christian")

	time.Sleep(10 * time.Second)
}

func talk(name string) {
	for {
		fmt.Printf("my name is %v\n", name)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}
