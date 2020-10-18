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
	//max should NOT be included
	max--
	//sum of 3 + 6 + 9 + ... an
	a1 := 3
	an := (max/a1) * a1
	n := an/a1
	sum := (a1 + an)*n/2
	//sum of 5 + 10 + 15 + ... + an
	a1 = 5
	an = (max/a1) * a1
	n = an/a1
	sum += (a1 + an)*n/2
	//sum of 3*5 = 15 + 30 + ... + an, which should be subtracted
	a1 = 15
	an = (max/a1) * a1
	n = an/a1
	sum -= (a1 + an)*n/2
	return sum
}

//func sum2(max, a int) int {
//	max
//}