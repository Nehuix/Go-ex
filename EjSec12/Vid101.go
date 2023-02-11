package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person // no se le da un nombre de campo porque es un struct
	ltk    bool
}

func main() {
	p1 := secretAgent{
		person: person{
			first: "Nehuen",
			last:  "Farias",
			age:   29,
		},
		ltk: true,
	}

	p2 := person{
		first: "Juli",
		last:  "Beola",
		age:   30,
	}

	fmt.Println(p1.first, p1.last, p1.age, p1.ltk)
	fmt.Println(p2)

	fmt.Println(p2.first, p1.last)
	fmt.Println(p2.last, p2.age)
}
