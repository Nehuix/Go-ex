package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	lado float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s *square) area() float64 {
	return s.lado * s.lado
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(&c)

	s := square{2}
	info(&s)
}
