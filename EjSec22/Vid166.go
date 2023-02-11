package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)

	fmt.Println("por salir.")
}

//receive from it
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Del canal eve:", v)

		case v := <-o:
			fmt.Println("Del canal odd:", v)

		case v := <-q:
			fmt.Println("Del canal quit:", v)
			return
		}
	}
}

//send to it
func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	q <- 0
}
