package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
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
	log.Info("STOP")
}

func PrimeFator(n int) map[int]bool {
	pf := map[int]bool{}
	return pf
}
