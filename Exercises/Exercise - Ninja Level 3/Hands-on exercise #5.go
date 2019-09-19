package main

import (
	"fmt"
)

func main() {
	for a := 11; a < 100; a++ {
		fmt.Printf("%v divided by 4: remainder is %v\n", a, a%4)
	}
}
