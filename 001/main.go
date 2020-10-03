package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"Imię"`
	Age   int    `json:"Wiek"`
}

func main() {
	var data = `
[
	{"Imię":"Paweł","Wiek":42},
	{"Imię":"Agnieszka","Wiek":40},
	{"Imię":"Lena","Wiek":7}
]`
	people := []Person{
		{
			First: "Paweł",
			Age:   42,
		},
		{
			First: "Agnieszka",
			Age:   40,
		},
	}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	err = json.Unmarshal([]byte(data), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
