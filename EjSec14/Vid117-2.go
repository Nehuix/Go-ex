package main

import (
	"fmt"
)

func main() {
	fmt.Println(bar()())

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
