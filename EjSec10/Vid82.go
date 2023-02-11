package main

import (
	"fmt"
)

// arrays no se usan, se usan slices
func main() {
	//x := type {values} composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x[1:3]) //no incluye el 3

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

}
