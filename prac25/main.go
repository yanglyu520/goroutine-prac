package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func() {
		for i := range 100 {
			ch1 <- strconv.Itoa(i)
		}
		ch2 <- 1
	}()

	go func() {
		for j := range ch1 {
			fmt.Println(j)
		}
		ch2 <- 2
	}()

	<-ch2
}
