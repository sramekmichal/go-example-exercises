package main

import "fmt"

func main() {
	a := struct {
		first_name       string
		friends          map[string]int
		favourite_drinks []string
	}{
		first_name: "Mike",
		friends: map[string]int{
			"John":    55,
			"Maribel": 29,
			"Mr. DJ":  32,
			"George":  46,
		},
		favourite_drinks: []string{
			"Pinacolada",
			"Water",
			"Coke",
		},
	}
	fmt.Println(a.first_name)
	fmt.Println(a.friends)
	fmt.Println(a.favourite_drinks)

	for k, v := range a.friends {
		fmt.Println(k, v)
	}

	for k2, v2 := range a.favourite_drinks {
		fmt.Println(k2, v2)
	}
}
