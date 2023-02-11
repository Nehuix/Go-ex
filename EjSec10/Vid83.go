package main

import (
	"fmt"
)

// arrays no se usan, se usan slices
func main() {
	//x := type {values} composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	//meto en  x los valores:...
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

}
