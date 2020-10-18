package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"sort"
	"strconv"
	"strings"
)

const taskDescr = `
--------------------------------------------------------------------------------
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
--------------------------------------------------------------------------------
`

type NumberE003 struct {
	Value         int
	LargestFactor int
	Factors       map[int]int
	Lim           int
}

func NewNumber(n int) *NumberE003 {
	t := NumberE003{}
	t.Value = n
	t.Lim = int(math.Sqrt(float64(n))) + 1
	t.Factors = map[int]int{}
	return &t
}

func (number *NumberE003) Factorize() {
	n := number.Value
	for {
		if n%2 == 0 {
			number.Factors[2]++
			if 2 > number.LargestFactor {
				number.LargestFactor = 2
			}
			n = n / 2
		} else {
			break
		}
	}
	for i := 3; i < number.Lim && i <= n; i += 2 {
		if n%i == 0 {
			number.Factors[i]++
			if i > number.LargestFactor {
				number.LargestFactor = i
			}
			n = n / i
		}
	}
}

func (number *NumberE003) String() string {
	factors := make([]string, len(number.Factors))
	factorsI := []int{}

	for key := range number.Factors {
		factorsI = append(factorsI, key)
	}
	sort.Ints(factorsI)
	for i, v := range factorsI {
		factors[i] = strconv.Itoa(v)
	}

	s := fmt.Sprintf("Value: %d. LargestFactor: %d. Factors: %v.", number.Value, number.LargestFactor, strings.Join(factors, ", "))
	return s
}

func main() {
	log.Info("START")
	log.Info("task description")
	log.Info(taskDescr)
	var number int
	flag.IntVar(&number, "number", 13195, "number to factorize")
	flag.Parse()
	n := NewNumber(number)

	n.Factorize()
	log.Info(n)
	log.Info("STOP")
}
