package main

import "fmt"

type person struct {
}
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Cześć")
}
func main() {
	p := person{}
	p.speak()
	say(&p)
}

func say(h human) {
	h.speak()
}
