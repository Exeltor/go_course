package main

import (
	"fmt"
	"math"
)

// create type square
// create type circle
// attach method to each that calculates area and returns it
// create type shape which defines an interface as anything which has the area method
// create func info which takes the type shape and prints area
// create value of type square
// create value of type circle
// use func info to print the area of the square
// use func info to print the area of the circle

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func main() {
	c := circle{
		radius: 6,
	}

	s := square{
		length: 5,
	}

	fmt.Println(info(c))
	fmt.Println(info(s))
}

func info(s shape) float64 {
	return s.area()
}
