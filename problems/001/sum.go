package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
)

const taskDescr = `If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The SimpleSum of these multiples is 23.
Find the SimpleSum of all the multiples of 3 or 5 below 1000.`

func main() {
	log.Info("START")
	log.Info(taskDescr)
	ch := make(chan int)
	var max int
	flag.IntVar(&max, "max", 10, "max value")
	flag.Parse()
	log.Infof("Searching for simple sum ov values under %d", max)
	go func() {
		ch <- SimpleSum(max)
	}()
	log.Infof("Searched value is %d", <-ch)
	log.Infof("Searching for arithmethic sum ov values under %d", max)
	go func() {
		ch <- ArithmeticSum(max)
	}()
	log.Infof("Searched value is %d", <-ch)

	log.Info("STOP")
}

func SimpleSum(max int) int {
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

func ArithmeticSum(max int) int {
	sum := sum2(max, 3)
	sum += sum2(max, 5)
	sum -= sum2(max, 15)
	return sum
}

func sum2(max, a int) int {
	if a > max {
		return 0
	}
	//max should NOT be included
	max--
	//sum of a1 + a2 + a3 + ... an
	a1 := a
	an := (max / a1) * a1
	n := an / a1
	return (a1 + an) * n / 2
}
