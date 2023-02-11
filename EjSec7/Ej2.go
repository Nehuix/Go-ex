package main

import (
	"fmt"
)

func main() {
	//for init; condition; post{}
	for i := 0; i <= 10; i++ {
		fmt.Printf("i es: %d\t\n", i)
		for j := 0; j < 3; j++{
			fmt.Printf("\t\t j es: %d\t\n", j)
		}
	}
}
