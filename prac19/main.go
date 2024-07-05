package main

import (
	"fmt"
	"sync"
)

func main() {
	var rwm sync.RWMutex
	var wg sync.WaitGroup
	m := make(map[int]int)
	wg.Add(5)

	for i := range 5 {
		go func(j int) {
			rwm.Lock()
			m[j] = j
			fmt.Println(m)
			rwm.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
