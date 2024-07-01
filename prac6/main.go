package main

import "fmt"

func main() {
	// make a nil channel of type string
	// A nil channel is not ready to send or receive data,
	// and any operation on it will block indefinitely.
	var ch1 chan string

	// unbuffered channel
	ch2 := make(chan int)
	ch4 := make(chan int, 0)

	// buffered channel with capacity of 2
	ch3 := make(chan string, 2)

	// a sent statement
	ch3 <- "go"
	ch2 <- 10

	// a receive expresion
	x := <-ch3

	// a receive statement; result is discarded
	// <-ch3

	// close channel will set program to panic
	close(ch3)
	fmt.Println(ch1, ch2, ch3, ch4, x)
}
