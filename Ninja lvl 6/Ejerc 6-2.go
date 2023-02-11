package main

import (
	"fmt"
)

func main() {
	vals := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(vals...))
	vals2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(bar(vals2))
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
