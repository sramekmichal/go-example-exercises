package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person1 struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person1{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {																	// Added error check
		log.Fatalln("JSON did not marshal - here is the error:", err)				// Chosen log function
	}
	fmt.Println(string(bs))

}
