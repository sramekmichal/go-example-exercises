package main

import "fmt"

func main() {
	if 150 == 150 {
		fmt.Println("true")
	}

	// Another type:

	x := "BUM"

	if x == "BUM" {
		fmt.Println(x)
	}
}