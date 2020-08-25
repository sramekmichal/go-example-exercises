package main

import(
	"fmt"
)

const(
	// Typed constants:
	a int = 486
	b float64 = 23.65
	c string = "I am happy"

	// Untyped constants:
	d = 487
	e = 23.66
	f = `"Bullshit"`
)

func main ()	{
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t", a, b, c, d, e, f)	// VALUE in the default format, in one line, TAb between each other
}