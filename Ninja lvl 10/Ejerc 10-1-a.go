package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) // buffered channel: le decis
	//cuantos vas a meter asi sabe
	c <- 42 //cuándo parar

	fmt.Println(<-c)
}
