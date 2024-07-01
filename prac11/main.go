package main

import (
	"fmt"
	"time"
)

type signal struct{}

func work() {
	fmt.Println("worker work now...")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)

	go func() {
		fmt.Println("worker goroutine starts work ...")
		f()
		c <- signal{} // push signal into c stream
	}()

	return c
}

func main() {
	c := spawn(work)
	<-c
	fmt.Println("worker work done", c)
}
