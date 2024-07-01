package main

import "fmt"

var c = make(chan int) // no buffered channel, send and receive happens at the same time
var a string

func f() {
	a = "x"
	<-c
	fmt.Println("2")
}
func main() {
	fmt.Println("0")
	go f()
	fmt.Println("1")
	c <- 5
	fmt.Println(a)
}
