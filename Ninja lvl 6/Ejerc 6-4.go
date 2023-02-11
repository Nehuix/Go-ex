package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Soy", p.first, p.last)
}

func main() {
	p1 := person{
		first: "Nehuen",
		last:  "Farias",
		age:   29,
	}
	p1.speak()
}
