package main

import "fmt"

func main() {
	m := map[string][]string	{
		"bond_james":		[]string{`Shaken, not stirred`, "Martinis", "Women"},
		"moneypenny_miss":	[]string{`James Bond`, "Literature", `Computer Science`},
		"no_dr":			[]string{`Being evil`, `Ice cream`, "Sunsets"},
	}
	fmt.Println(m)

	m["vomacka_franta"] = []string{`You shall ask me`, "Food", "Cars"}		// Added a note to a map

	b := "\nFavourite things:"

	for key, value := range m {
		fmt.Println("\nThis is the record for: ", key, b)
		for index, value2 := range value {
			fmt.Println("\t\t\t", index, value2)
		}
	}
}