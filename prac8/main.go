package main

import "fmt"

// channels to make a pipeline
// first go routine generates 0,1,2..
// second goroutine squares each int
// third one prints each int

func main() {
	naturalstream := make(chan int)
	squarestream := make(chan int)

	go func() {
		for x := 0; x < 15; x++ {
			naturalstream <- x
		}
		close(naturalstream) // this will stop the next process to endlessly waiting for natural stream
	}()

	go func() {
		for x := range naturalstream {
			fmt.Println("x in naturalstream: ", x)
			squarestream <- x * x
		}
		close(squarestream)
	}()

	// third goroutine is the main goroutine
	// use a range key word, so that the channel is naturally closed when loop is end
	for x := range squarestream {
		fmt.Println("x in squarestream: ", x)
	}
}
