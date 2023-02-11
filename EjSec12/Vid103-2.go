package main

import (
	"fmt"
)



func main() {
	p1 := struct {
	first string
	last  string
	age   int
}{
		first: "Nehuen",
		last: "Farias",
		age: 29,
	}

	fmt.Println(p1)
}