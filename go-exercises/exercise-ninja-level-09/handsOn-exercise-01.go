package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin CPU:\t", runtime.NumCPU())
	fmt.Println("begin GoRout:\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("First go routine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second go routine")
		wg.Done()
	}()

	fmt.Println("mid CPU:\t", runtime.NumCPU())
	fmt.Println("mid GoRout:\t", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("EXIT")

	fmt.Println("end CPU:\t", runtime.NumCPU())
	fmt.Println("end GoRout:\t", runtime.NumGoroutine())

}
