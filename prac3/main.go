package main

import (
	"fmt"
	"time"
)

// because the main goroutine and the spinner goroutine is separate
// the spinner goroutine will stop when the goroutine stops
// if we only use spinner func, it will be an infinite loop
// if we set up a timer for spinner, we have to calculate the time that fib() function runs, which changes depends on the parameter passed in
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d \n", n, fibN)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
