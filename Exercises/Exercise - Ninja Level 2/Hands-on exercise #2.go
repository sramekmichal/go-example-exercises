package main

import (
	"fmt"
)

func main ()	{
	x := 23
	y := 15
	fmt.Println(x == y)
	fmt.Println(x <=y)
	fmt.Println(x >= y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x > y)
	fmt.Printf("\n")		// Only shows new line

	// Another method for the same result:

	a := (23 == 15)
	b := (23 <= 15)
	c := (23 >= 15)
	d := (23 != 15)
	e := (23 < 15)
	f := (23 > 15)
	fmt.Println(a, b, c, d, e, f)
}

