package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person2 struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person2{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
//		fmt.Println(err)				choice 1
//		return

//		log.Println(err)				choice 2
//		return

		log.Fatalln(err)				//choice 3
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
	}
	return bs, nil
}