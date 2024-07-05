package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	for i := range 5 {
		go func(j int) {
			ch1 <- j
		}(i)
	}

	for {
		select {
		case a1 := <-ch1:
			fmt.Println(a1)
		default:
			fmt.Println("default")
		}
	}

}
