package main

import (
	"fmt"
)
// Solution func literal a):
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// Solution buffered channel b):
func main() {
	c := make(chan int, 1)

		c <- 42

	fmt.Println(<-c)
}

// Solution b) is not ideal, but works.