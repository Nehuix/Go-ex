package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

var x int
func main() {
	p1 := person{
		first : "Nehuen",
		last: "Farias",
		age: 29,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	(*p).first = "Juli"  //(*p) es desrreferenciar. O sea te da el
						//valor de esa dirección
}