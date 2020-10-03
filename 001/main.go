package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Hello")
	x := `Pa1sword 1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(x), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(x)
}
