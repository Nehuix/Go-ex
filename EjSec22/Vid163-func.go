package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan <- int)
	cr := make(<- chan int)

	fmt.Println("----------------------------------------------------------------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)
}