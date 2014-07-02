package main

import "fmt"

type Rect struct {
	Width  int
	Height int
}

type Square struct {
	Side int
}

type Circle struct {
	Radius int
}

type Shape interface {
	Area() int
}

func main() {
	shapes := []Shape{
		&Rect{40, 53},
		&Square{33},
		&Circle{92},
	}

	for i, v := range shapes {
		fmt.Printf("shape %v area: %v\n", i, v.Area())
	}
}
