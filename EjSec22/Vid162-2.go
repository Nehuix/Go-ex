package main

import "fmt"

func main() {
	c := make(chan int, 2) // canal buffer, permite guardar x vlaores
						   // sin importar si hay algo listo para sacarlo
	c <- 42 //pone el 43 en c
	c <- 43 //no funciona porque queres poner mas de uno (si pongo dos puedo 2)
	
	fmt.Println(<-c)
	fmt.Println(<-c)

	// evitar buffer chanels
}