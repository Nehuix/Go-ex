package main

import (
	"fmt"
)

type person struct {
	first  string
	last   string
	favice []string
}

func main() {
	p1 := person{
		first: "Nehuen",
		last:  "Farias",
		favice: []string{
			"Limon",
			"dulc",
			"granizado",
		},
	}

	p2 := person{
		first: "Juli",
		last:  "Beola",
		favice: []string{
			"Banana",
			"dulcede",
			"menta",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)

	for i, v := range p1.favice {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	fmt.Println(p2.favice[2])

	for i, v := range p2.favice {
		fmt.Println(i, v)
	}

}
