package main

import (
	"fmt"
	"math"
)

type square struct {
	lado float64
}
type circle struct {
	radio float64
}

func (c circle) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (s square) area() float64 {
	return s.lado * s.lado
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circ := circle{radio: 12.345}

	squa := square{lado: 15}

	info(circ)
	info(squa)
}
