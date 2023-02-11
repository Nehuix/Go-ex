package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GRs:", runtime.NumGoroutine())

	var counter int64
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)


	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			//time.Sleep(time.Second)
			
			wg.Done()
		}()
		fmt.Println("GRs nuevass:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("GRs nuevas:", runtime.NumGoroutine())
	fmt.Println("Contador:", counter)
}
