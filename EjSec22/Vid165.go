package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	for v:=range c{
		fmt.Println(v)
	}
	fmt.Println("Por salir")
}

//send to it
func foo(c chan <-int) { //nueva area de memoria
	for i:=0;i<10;i++{
		c <- i
	}
	close(c) //esto comunica que el canal se completó y
			 //no se van a mandar mas valores en el	
}


