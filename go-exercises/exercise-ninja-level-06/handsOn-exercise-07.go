package main

import (
	"fmt"
)

func main() {

	f := func() {
		fmt.Println("Here is some anonymous func")
	}
	f()
	fmt.Printf("%T\n", f)

	fmt.Println("DONE")
}
