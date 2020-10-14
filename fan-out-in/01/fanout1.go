package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

const max_fib = 100

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	result := []int{}
	for v := range c2 {
		result = append(result, v)
	}
	sort.Ints(result)
	for idx, v := range result {
		fmt.Println(idx, v)
	}
	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < max_fib; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup
	goroutines := runtime.NumCPU()
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	//time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	var a, b, c int
	a += 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	for i := 0; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
