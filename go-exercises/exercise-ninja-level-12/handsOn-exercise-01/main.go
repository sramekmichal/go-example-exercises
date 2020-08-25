package main

import (
	"fmt"
	"github.com/sramek5/example/go-exercises/exercise-ninja-level-12/handsOn-exercise-01/dog"
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