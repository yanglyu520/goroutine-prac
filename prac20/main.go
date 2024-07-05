package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	time.AfterFunc(3e9, func() {
		fmt.Println("delayed func")
	})

	time.Sleep(4e9)
	fmt.Println("end")
}
