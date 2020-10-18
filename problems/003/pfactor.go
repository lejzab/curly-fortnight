package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"math"
	"sort"
)

const taskDescr = `The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

`

func main() {
	log.Info("START")
	log.Info(taskDescr)
	var number int
	flag.IntVar(&number, "number", 13195, "number to factorize")
	flag.Parse()
	pf := PrimeFator(number)
	pf1 := []int{}
	for v := range pf {
		pf1 = append(pf1, v)
	}
	sort.Ints(pf1)

	log.Infof("The prime factors of %d are:\n", number)
	for _, v := range pf1 {
		log.Infof("\t%d. Number of this factor: %d", v, pf[v])
	}
	log.Infof("The largest prime factor of %d is: %d", number, pf1[len(pf1)-1])
	log.Info("STOP")
}

func PrimeFator(n int) map[int]int {
	pf := map[int]int{}
	lim := int(math.Sqrt(float64(n))) + 1
	for {
		if n%2 == 0 {
			n = n / 2
			pf[2]++
		} else {
			break
		}
	}
	for i := 3; i < lim && i <= n; i += 2 {
		if n%i == 0 {
			n = n / i
			pf[i]++
		}
	}

	return pf
}
