package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)

	fmt.Println("por salir.")
}

//receive from it
func receive(eve, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-eve:
			fmt.Println("Del canal eve:", v)

		case v := <-odd:
			fmt.Println("Del canal odd:", v)

		case i, ok := <-quit: //check si está cerrado el canal
			if !ok{
				fmt.Println("from comma ok", i, ok)
			return
			} else {
				fmt.Println("from comma ok", i)
			}
			
		}
	}
}

//send to it
func send(eve, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}
