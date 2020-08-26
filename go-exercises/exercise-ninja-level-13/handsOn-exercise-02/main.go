package main

import (
	"fmt"
	"github.com/sramek5/example/go-exercises/exercise-ninja-level-13/handsOn-exercise-02/quote"
	"github.com/sramek5/example/go-exercises/exercise-ninja-level-13/handsOn-exercise-02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}