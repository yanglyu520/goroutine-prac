package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("in the separate goroutine")
		ch <- 998
	}()

	a, ok := <-ch
	fmt.Println(a, ok)
}
