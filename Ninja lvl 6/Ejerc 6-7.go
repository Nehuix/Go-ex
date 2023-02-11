package main

import (
	"fmt"
)

var x int
var g func()

func main() {

	a := foo()
	
	fmt.Println(a)

	b := func (){
		for i := 0; i < 10; i++ {
			fmt.Println("Num:", i)
		}
	}
	b()
	fmt.Printf("%T\n", b)

	fmt.Println("Ahora g: ")

	g = b
	g()

}

func foo() string{
	return "Tuvie"
}
