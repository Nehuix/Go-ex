package main

import (
	"fmt"
)

// arrays no se usan, se usan slices
func main() {
	//make(type, length, capacity)
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)
	x = append(x, 3425)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
