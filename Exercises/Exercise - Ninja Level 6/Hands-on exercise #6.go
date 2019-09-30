package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Here is some anonymous func")
	}()
}

func foo() {
	fmt.Println("sounds easy?")
}
