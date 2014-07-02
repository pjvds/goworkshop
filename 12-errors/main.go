package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		n := GetNumber()
		fmt.Printf("number: %v", n)
	}
}

func GetNumber() (int, error) {
	if n := rand.Intn(100); n%2 == 0 {
		return n, nil
	}
	return 0, errors.New("connection error")
}
