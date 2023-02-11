package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)){code}
// cada valor del tipo secretAgent tiene acceso al metodo
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Nehuen",
			"Farias",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
