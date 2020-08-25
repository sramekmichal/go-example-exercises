package main

import (
	"fmt"
)

func main() {
	x := 1991
	for {
		if x > 2020 {
			break
		}
		fmt.Println(x)
		x++
	}
}