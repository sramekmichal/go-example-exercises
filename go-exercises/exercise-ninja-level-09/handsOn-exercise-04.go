package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var m sync.Mutex

func main() {
	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			m.Unlock()
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
