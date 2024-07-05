package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("carry out the %dth execution\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("done in main")
}
