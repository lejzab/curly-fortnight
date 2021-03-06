package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	//var mu sync.Mutex

	for i := 0.; i < gs; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("COUNTER", counter)
}
