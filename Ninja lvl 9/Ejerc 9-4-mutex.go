package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	counter :=0
	gs:=100
	wg.Add(gs)

	for i:=0; i<gs; i++{
		go func(){
			mu.Lock()
			v:=counter
			//runtime.Gosched() lo sacas porque no tiene sentido pausar todo si usas lock
			v++
			counter=v
			fmt.Println("counter:", counter)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Valor final:", counter)
}