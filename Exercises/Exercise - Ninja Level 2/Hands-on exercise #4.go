package main

import(
	"fmt"
)

func main()		{
	a := 211
	fmt.Printf("%d\t\t%b\t\t%#x\n", a, a, a)

	// Shift the bits over 1 position to the left
	y := a << 1
	fmt.Printf("%d\t\t%b\t\t%#x", y, y, y)

}