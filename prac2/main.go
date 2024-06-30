package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Set the number of CPUs to use

	var wg sync.WaitGroup
	wg.Add(4)

	for i := 0; i < 4; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				fmt.Printf("Goroutine %d - iteration %d\n", id, j)
			}
		}(i)
	}

	wg.Wait()
}
