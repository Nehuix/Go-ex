package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	bar(c) //sin el go se bloquea ahi hasta 
		   //que haga lo que tiene que hacer
	fmt.Println("Por salir")
}

//send to it
func foo(c chan <-int) { //voy de general a específico
	c <- 42
}

//receive from it
func bar(c <-chan int){
	fmt.Println(<-c)
}