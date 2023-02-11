package main

import (
	"fmt"
)

const (
	_ = iota
	//kb := 1024
	kb = 1 << (iota * 10) //iota=1
	mb = 1 << (iota * 10) //iota=2
	gb = 1 << (iota * 10) //iota=3
)

func main() {

	fmt.Printf("%d\t\t\t\t%b\n", 1, 1)
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
