package main

import(
	"fmt"
)
type chleba int
var x chleba

func main()	{
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}