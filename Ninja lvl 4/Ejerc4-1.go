package main

import (
	"fmt"
)

func main() {
	x := [5]int{1, 54, 87, 2, 6}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}
