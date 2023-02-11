package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42 //pone el 43 en c
	}()
	
	fmt.Println(<-c)
}