package main

import "fmt"

type Rect struct {
	Width  int
	Height int
}

func main() {
	r := Rect{
		Width:  88,
		Height: 370,
	}

	fmt.Printf("area: %v\n", r.Width*r.Height)
}
