package main

import (
	"fmt"
)

func main() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(x1)
	fmt.Println(x2)

	xxs := [][]string{x1, x2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("Record: ", i)
		for j, v := range xs {
			fmt.Printf("\tPosicion: %v\t valor: \t%v", j, v)
			fmt.Println("")
		}
	}
}
