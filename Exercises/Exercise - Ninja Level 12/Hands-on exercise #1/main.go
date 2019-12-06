package main

import (
	"fmt"
	"github.com/sramek5/example/Exercises/Exercise - Ninja Level 12/Hands-on exercise #1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	Husky := canine {
		name: "Caesar",
		age: dog.Years(10),
	}
	fmt.Println(Husky)
}
