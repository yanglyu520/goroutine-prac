package main

import "fmt"

// channels to make a pipeline
// first go routine generates 0,1,2..
// second goroutine squares each int
// third one prints each int

func counter(naturalsInStream chan<- int) {
	for x := 0; x < 15; x++ {
		naturalsInStream <- x
	}
	close(naturalsInStream)
}

func squarer(naturalsOutStream <-chan int, squareInStream chan<- int) {
	for x := range naturalsOutStream {
		fmt.Println("x in naturalstream: ", x)
		squareInStream <- x * x
	}
	close(squareInStream)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	naturalsStream := make(chan int)
	squarestream := make(chan int)

	go counter(naturalsStream)
	go squarer(naturalsStream, squarestream)
	printer(squarestream)
}
