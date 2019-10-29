package main

import (
	"fmt"
)

type person struct {
	first string
}

func (p person) speak () {
	fmt.Println("It works!")
}

type human interface {
	speak()
}

func saySomething (h human){
	h.speak()
}

func main(){
	p1 := person{
		first: "Vlado",
	}
	saySomething(&p1)				// Can pass a value of type *person
//	saySomething(p1)				// Cannot pass a value of type person

	p1.speak()
}


