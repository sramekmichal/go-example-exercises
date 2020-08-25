package main

import "fmt"

func main() {
	x := 500
	if x > 500 {
		fmt.Println("Fatal Error")
	} else if x < 500 {
		fmt.Println("Wrong Way")
	} else {
		fmt.Println(x)
	}
}