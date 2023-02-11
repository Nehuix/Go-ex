package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}


func main() {
	p1 := person {
		first: "Nehuen",
		last: "Farias",
		age: 32,
	}

	fmt.Println(p1)
}