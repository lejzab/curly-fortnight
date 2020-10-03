package main

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
	"sort"
)

func main() {
	log.Info("START")
	xi := []int{7, 10, 1, 2, 6, 5, 12}
	log.Info(xi)
	sort.Ints(xi)
	log.Info(xi)
	xs := []string{"Paweł", "Lena", "Agnieszka", "Aurelia", "Jagoda", "Kuba", "Łukasz", "Jeremi"}
	log.Info(xs)
	coll := collate.New(language.Polish)
	coll.SortStrings(xs)
	log.Info(xs)
	log.Info("STOP")
}
