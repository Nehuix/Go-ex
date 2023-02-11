package main

import (
	"fmt"
)

func main() {
	a := func(xi []int) int {
		if len(xi)==0{return 0}
		
		if len(xi)==1{return 1}

		return xi[0] + xi[len(xi)-1]
	}

	x := foo(a, []int{1, 2, 3,4,8,101})
	fmt.Println(x)
}

func foo(f func (xi []int) int, ii []int) int{
	n := f(ii)
	fmt.Println(n)
	n++
	return n
}
