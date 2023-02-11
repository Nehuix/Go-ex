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
	p1 := person{
		first: "Nehuen",
		last:  "Farias",
		age:   29,
	}

	p2 := person{
		first: "Juli",
		last:  "Beola",
		age:   30,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p2.first, p1.last)
	fmt.Println(p2.last, p2.age)
}
