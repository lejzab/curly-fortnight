package main

import (
	"fmt"
)

func doSome(x int) int {
	return x * x
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSome(12)
	}()
	fmt.Println(<-ch)
}
