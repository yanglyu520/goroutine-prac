package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once

	wg.Add(10)

	for i := range 10 {
		go func() {
			defer wg.Done()
			fmt.Println(i)
			once.Do(func() {
				fmt.Println("heavy goroutine")
				time.Sleep(1e9)
			})
		}()
	}

	wg.Wait()
}
