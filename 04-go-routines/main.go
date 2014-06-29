package main

import (
	"fmt"
	"time"
)

func main() {
	talk("rob")
	talk("christian")
}

func talk(name string) {
	for {
		fmt.Printf("my name is %v\n", name)
		time.Sleep(1 * time.Second)
	}
}
