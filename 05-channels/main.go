package main

import (
	"fmt"
	"time"
)

func main() {
	go talk("rob")
	go talk("christian")

	time.Sleep(10 * time.Second)
	fmt.Println("I'm tired of you talking,.. bye!")
}

func talk(name string) {
	for {
		fmt.Printf("my name is %v\n", name)
		time.Sleep(1 * time.Second)
	}
}
