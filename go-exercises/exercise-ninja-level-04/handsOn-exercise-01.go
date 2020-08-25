package main

import "fmt"

func main ()	{
	x := [5]int{5, 12, 36, 5000, 2}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}