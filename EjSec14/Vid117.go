package main

import (
	"fmt"
)

func main() {
	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)

}

func foo() string {
	s := "Hola mundo"
	return s
}

// func    return - lo que retorna
func bar() func() int {
	return func() int {
		return 451
	}
}
