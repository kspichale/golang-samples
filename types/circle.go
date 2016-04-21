package main

import (
  "fmt"
  "math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := Circle{2}
	fmt.Println(c.Area())
}
