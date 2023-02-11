package main

import (
	"fmt"
)

type person struct {
	Nombre   string
	Apellido string
	Edad     int
	ColFav   string
}

type human interface {
	speak() // se le pondría el tipo despues si retornara algo
}

func (p person) speak() {
	fmt.Println("Hola, soy", p.Nombre, p.Apellido)
	fmt.Println("Tengo", p.Edad, "y mi color favorito es", p.ColFav)
}

func main() {
	p1 := person{
		Nombre:   "Nehuen",
		Apellido: "Farias",
		Edad:     29,
		ColFav:   "Rojo",
	}

	p2 := person{
		Nombre:   "Juli",
		Apellido: "Beola",
		Edad:     29,
		ColFav:   "Verde?",
	}

	saySomething(&p1)
	saySomething(&p2)
}

func saySomething(h human) {
	h.speak()
}
