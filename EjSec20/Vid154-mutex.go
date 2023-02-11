package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GRs:", runtime.NumGoroutine())

	counter := 0
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() //nadie puede puede leer esto hasta que termine
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GRs nuevass:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("GRs nuevas:", runtime.NumGoroutine())
	fmt.Println("Contador:", counter)
}
