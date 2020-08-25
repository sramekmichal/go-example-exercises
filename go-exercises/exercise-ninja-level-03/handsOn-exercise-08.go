package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("FU")
	case true:
		fmt.Println(`"Hello"`)
	}
}
