package main

import (
	"fmt"
	"github.com/sramek5/example/go-exercises/exercise-ninja-level-13/handsOn-exercise-01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(10))
}