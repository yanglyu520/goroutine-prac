package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func() {
		ch1 <- "first data"
		ch2 <- 1
	}()
	go func() {
		content := <-ch1
		fmt.Println(content)
		ch2 <- 2
	}()

	<-ch2
	<-ch2

	fmt.Println("done in main goroutine")
}
