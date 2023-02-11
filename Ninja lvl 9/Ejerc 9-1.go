package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("CPU inicio:", runtime.NumCPU())
	fmt.Println("GRs inicio:", runtime.NumGoroutine())
	wg.Add(3)

	go goR1()
	go goR2()
	go goR3()
	fmt.Println("CPU mitad:", runtime.NumCPU())
	fmt.Println("GRs mitad:", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Por finalizar")
	fmt.Println("CPU fin:", runtime.NumCPU())
	fmt.Println("GRs fin:", runtime.NumGoroutine())

}

func goR1() {
	fmt.Println("Esta es la goR1")
	for i := 0; i < 10; i++ {
		fmt.Println("Cosita:", i)
	}
	wg.Done()
}

func goR2() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Núm:", i)
	}
}

func goR3() {
	fmt.Println("Ahora viene la goR3")
	wg.Done()
}
