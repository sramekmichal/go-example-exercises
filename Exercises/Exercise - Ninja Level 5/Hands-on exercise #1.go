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

	fmt.Println(x1.first_name, x1.last_name)

	for k, v := range x1.favourite_ice_cream_flavor {
		fmt.Println(k, v)
		}

	fmt.Println(x2.first_name, x2.last_name)
	for k, v := range x2.favourite_ice_cream_flavor {
		fmt.Println(k, v)
		}
}