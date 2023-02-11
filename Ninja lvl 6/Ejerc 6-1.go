package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	x, s := bar()
	fmt.Println(x)
	fmt.Println(s)

}

func foo() int {
	return 3
}

func bar() (int, string) {
	return 5, "hola"
}
