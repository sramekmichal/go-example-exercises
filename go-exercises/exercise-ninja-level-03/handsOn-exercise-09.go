package main

import "fmt"

func main() {
	favSport := "chess"
	switch favSport {
	case "bicycle":
		fmt.Println("boring")
	case "airsoft":
		fmt.Println("kinda fun")
	case "chess":
		fmt.Println("No. 1")
	case "JackAss":
		fmt.Println(`"Do not try this at home"`)
	}
}
