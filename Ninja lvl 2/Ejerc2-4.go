package main

import (
	"fmt"
)

func main() {
	b := 50

	fmt.Printf("%d\t%b\t%x", b, b, b)
	a := b << 1
	fmt.Println("")
	fmt.Printf("%d\t%b\t%x", a, a, a)

}
