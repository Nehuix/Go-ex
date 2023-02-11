package main

import (
	"fmt"
)

type jeje string
var x jeje
var y string

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = "holiis"
	fmt.Println(x)
	y = string(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}