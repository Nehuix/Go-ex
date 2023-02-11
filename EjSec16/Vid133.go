package main

import (
	"fmt"
)

func main() {
	x:=0
	fmt.Println("La dire antes",&x)
	fmt.Println("Antes", x)
	foo(&x)
	fmt.Println("Despues",&x)
	fmt.Println("Despues",x)
}

func foo(y *int){
	fmt.Println("y antes",y)
	fmt.Println("y antes",*y)
	*y=43
	fmt.Println("y despues",y)
	fmt.Println("y despues",*y)
}