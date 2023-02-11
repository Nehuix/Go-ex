package main

import (
	"fmt"
)

func main() {
	foo()

	f := func() {
		fmt.Println("Func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("Año:", x)
	}
	g(1984)

}

func foo() {
	fmt.Println("hello from foo")
}
