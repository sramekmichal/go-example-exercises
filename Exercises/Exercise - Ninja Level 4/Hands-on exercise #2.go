package main

import "fmt"

func main ()	{
	x := []int{2, 4, 25, 12, 585, 46, 7, 12, 3, 66}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}