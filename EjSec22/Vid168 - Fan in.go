package main

import (
	"fmt"
	"sync"
)

//Fanin es un canal donde se guardan los resultados de
//varios canales para despues sacarlos juntos
func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive (even, odd, fanin)

	for v := range fanin{
		fmt.Println(v)
	}
}

//send channel
func send(even, odd chan <- int){ //guarda todos los pares en
	for i:=0; i<10; i++{		  //el canal de pares e impares
		if i%2 ==0{				  //en el de impares
			even <- i
		} else {
			odd <- i
		}
	}
	close (even)
	close (odd)
}

//receive channel
func receive(even, odd <- chan int, fanin chan <- int){
	var wg sync.WaitGroup
	wg.Add(2)

	go func (){  			//primero saca los numeros del canal
		for v:= range even{ //par y los guarda en el fanin
			fanin <-v		//y despues hace lo mismo con los impares
		}
		wg.Done()
	}()

	go func (){
		for v := range odd{
			fanin <-v
		}
		wg.Done()
	}()

	wg.Wait()			   //espera a que terminen de guardarse
	close(fanin)
	fmt.Println("Terminó el canal")

}