package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := circle{2}
	fmt.Println(c.Area())
}
