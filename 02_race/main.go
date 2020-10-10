package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	//wg.Add(gs)

	for i := 0.; i < gs; i++ {
		wg.Add(1)
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("COUNTER", counter)
}
