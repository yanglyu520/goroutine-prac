package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	for i := 0; i < 5; i++ {
		fmt.Printf("round %d: \n", i)
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println("data from ch1: ", data)
			} else {
				fmt.Println("ch1 is closed")
			}
		case data, ok := <-ch2:
			if ok {
				fmt.Println("data from ch2: ", data)
			} else {
				fmt.Println("ch2 is closed")
			}
		default:
			fmt.Println("default")
		}
	}

}
