package main

import (
	"fmt"
)

func main() {
	a := retFun()
	
	fmt.Println("Es", a(67))
}

func retFun () func (i int) int{
	return func (i int) int {
		return i+1
	}
}