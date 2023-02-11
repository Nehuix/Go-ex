package main

import "fmt"

func main() {
	c := make(<- chan int, 2) //recieve only
						      //el <- chan hace que solo se puedan
							  //recibir datos al channel
	c <- 42 //pone el 43 en c
	c <- 43 //no funciona porque queres poner mas de uno (si pongo dos puedo 2)
	
	fmt.Println(<-c)
	fmt.Println(<-c)
	// evitar buffer chanels

	fmt.Println("----------------------------------------------------------------")
	fmt.Printf("%T\n", c)
}