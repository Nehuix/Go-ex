package main

import (
	"fmt"
)

func main() {
	x := []int{1, 54, 87, 2, 6, 47, 8, 4, 11, 10}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}
