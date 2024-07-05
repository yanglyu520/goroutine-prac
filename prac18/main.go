package main

import (
	"fmt"
	"sync"
)

var (
	num = 10
	wg  sync.WaitGroup
	m   sync.Mutex
)

func main() {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			m.Lock()
			num--
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("done in main", num)
}
