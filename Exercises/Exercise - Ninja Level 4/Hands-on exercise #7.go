package main

import "fmt"

func main() {
	words := []string{"James", "Bond", "Shaken, not stirred"}
	sentences := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	fmt.Println(words)
	fmt.Println(sentences)

	quotes := [][]string{words, sentences}
	fmt.Println(quotes)

	for i, c := range quotes {
		fmt.Println("record: ", i)
		for j, val := range c {
			fmt.Printf("\tindex posiiton: %v\t value \t%v\n", j, val)
		}
	}
}