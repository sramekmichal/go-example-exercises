package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
}

type CIRCLE struct {
	radius float64
}

func (s SQUARE) area() float64 {
	return s.length * s.length
}

func (c CIRCLE) area() float64 {
	return math.Pi * c.radius * c.radius
}

type SHAPE interface {
	area() float64
}

func INFO (s SHAPE) {
	fmt.Println(s.area())
}

func main() {
	circ := CIRCLE{
		radius: 12.345,
		}
	squa := SQUARE{
		length: 15,
	}
	fmt.Printf("circle area: ")
		INFO(circ)
	fmt.Printf("square area: ")
		INFO(squa)
}

