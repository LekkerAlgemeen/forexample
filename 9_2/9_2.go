package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   37,
	}

	saySomething(&p)
	p.speak()
}

func (p *person) speak() {
	fmt.Println("Hi, my name is", p.first, p.last, "and my age is", p.age)
}

func saySomething(h human) {
	h.speak()
}
