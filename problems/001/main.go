package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const taskDescr = `If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The Sum of these multiples is 23.
Find the Sum of all the multiples of 3 or 5 below 1000.`

func main() {
	fmt.Println(taskDescr)
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Sum(max))
}

func Sum(max int) int {
	v := map[int]bool{}
	for i := 3; i < max; i += 3 {
		v[i] = true
	}
	for i := 5; i < max; i += 5 {
		v[i] = true
	}
	sum := 0
	for k := range v {
		sum += k
	}
	return sum
}
