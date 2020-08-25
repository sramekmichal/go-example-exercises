package main
import (
	"fmt"
)
type person struct {
	first_name string
	last_name string
	favourite_ice_cream_flavor []string
}
func main() {
	x1 := person{
		first_name:                 "Ace",
		last_name:                  "Ventura",
		favourite_ice_cream_flavor: []string{"facebook blue", "apricot", "strawberry"},
	}
	x2 := person{
		first_name:                 "Holy",
		last_name:                  "Shit",
		favourite_ice_cream_flavor: []string{"brown", "mellon", "cherry"},
	}
	m := map[string]person{
		x1.last_name: x1,
		x2.last_name: x2,
	}
		for _, v := range m {
		fmt.Println(v.first_name, v.last_name)
				for i, v2 := range v.favourite_ice_cream_flavor {
					fmt.Println(i, v2)
				}
				fmt.Println("*--------------*")
		}
	}