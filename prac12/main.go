package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct{}

func worker(i int) {
	fmt.Printf("worker %d starts now...\n", i)
	time.Sleep(1 * time.Second)
}

func spawn(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			<-groupSignal
			fmt.Println("workers goroutine starts work ...")
			f(i)
			wg.Done()
		}()
		go func() {
			wg.Wait()
			c <- signal{}
		}()
		return c
	}

	return c
}

func main() {
	fmt.Println("start a group of workers")
	groupSignal := make(chan signal)
	c := spawn(worker, 5, groupSignal)
	time.Sleep(5 * time.Second)
	close(groupSignal)
	<-c
	fmt.Println("work done")
}
