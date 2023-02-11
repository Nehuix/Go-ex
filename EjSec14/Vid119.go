package main

import (
	"fmt"
)

func main() {
	a:= incrementor() //incrementor es la func a la que llamada
	b:= incrementor() // y retorna una función. Si fuese una variable cualquiera
	fmt.Println(a())  // seria como si guardara un valor cualñquiera
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int{
	var x int
	fmt.Println("Increm")

	return func() int{
		x++
		return x
	}
}
