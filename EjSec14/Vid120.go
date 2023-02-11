package main

import (
	"fmt"
)

func main() {
	n:=factorial(4)
	fmt.Println(n)
	r:=loopFact(4)
	fmt.Println(r)		
}

func factorial(n int) int{
	if n==0{
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int{
	total := 1
	for ; n>0; n--{ // no tiene iniciador por eso es for "nada";...
		total*=n
	}
	return total
}