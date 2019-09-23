package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct{
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}
func main() {
	c1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "green",
		},
		fourWheel: true,
	}

	fmt.Println(c1)

	c2 := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(c2)
	fmt.Printf("%v\n%v\n", c1.color, c2.color)
}
