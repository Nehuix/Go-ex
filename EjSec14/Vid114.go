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
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface { //cualquiera que tenga un tipo que tenga el
	speak() //metodo speak(), tambien es de tipo humano
} //un valor puede ser de más de un tipo

func bar(h human) {
	//fmt.Println("Me pasaron a bar", h)
	switch h.(type) { //switchea el type
	case person:
		fmt.Println("Me pasaron a barrrr", h.(person).first)
	case secretAgent:
		fmt.Println("Me pasaron a barrr", h.(secretAgent).first)
	}
	fmt.Println("Me pasaron a bar", h)
}

type pancho int

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

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
	bar(sa1)
	bar(sa2)
	bar(p1)

	//conversion
	var x pancho = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
