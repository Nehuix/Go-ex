package main

import (
	"fmt"
)

func main() {
	defer foo()
	fmt.Println("Holis mundis 1")
}

func foo() {
	defer func() {
		fmt.Println("Defer de foo 3")
	}()
	fmt.Println("foo pri 2")
}
