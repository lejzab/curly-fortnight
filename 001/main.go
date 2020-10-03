package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Age   int
}

var data = `
[
	{"First":"Paweł","Age":42},
	{"First":"Agnieszka","Age":40},
	{"First":"Lena","Age":7}
]`

func main() {
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
