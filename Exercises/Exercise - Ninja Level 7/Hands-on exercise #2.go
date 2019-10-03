package main

import "fmt"

type person struct {
	first string
	last string
}

func main() {
	p1 := person {
		first: "Hans",
		last: "Jablonski",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	// Method ia):
	p.first = "Vaclav"
	p.last = "Klaussen"

	// Method ib):
	(*p).first = "Vaclavik"
	(*p).last = "Claus"
}
