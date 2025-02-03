package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println(multiArea(&c, &r))
}

// declaration of interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x, y, l, b float64
}

// to implement an interface for particular struct implement its methods
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	return r.l * r.b
}

func multiArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.area()
	}
	return totalArea
}
