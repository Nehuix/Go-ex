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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println("El mapa es:")

	for _, v := range m {
		fmt.Println(v.first)
		for i, val := range v.favice {
			fmt.Println(i, val)
		}
	}

}
