package main

import (
	"fmt"
)

const (
	y1 = 2020
	y2 = iota + y1
	y3 = iota + y1
	y4 = iota + y1

	// Another possible solution:
	x1 = 2016 + iota
	x2 = 2016 + iota
	x3 = 2016 + iota
	x4 = 2016 + iota
)

func main ()	{
	fmt.Println(y1, y2, y3, y4)
	fmt.Println(x1, x2, x3, x4)
}
