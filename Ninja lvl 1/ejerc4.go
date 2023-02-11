package main

import (
	"fmt"
)
type jeje int
var x jeje
func main(){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}