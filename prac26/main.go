package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	for i := range 5 {
		go func(j int) {
			fmt.Println("goroutine", j)
			ch <- j
		}(i)
	}

	for k := range 5 {
		x := <-ch
		fmt.Println("main", k, x)
		time.Sleep(1e9)

	}
}
