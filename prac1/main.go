package main

import (
	"fmt"
	"time"
)

// 实际上，计算机的处理器通常会使用一种名为分时的技术，在多个goroutine 上面轮流 花费 一些时间。因为分时的具体实施细节通常只有Go运行时、操作系统和处理器会知道， 所以我们在使用goroutine的时候，应该假设不同goroutine中的各项操作将以任意顺序执行
func main() {
	for i := range 4 {
		go sleepyGopher(i)
	}

	time.Sleep(4 * time.Second)
	fmt.Println("wake up...")
}

func sleepyGopher(i int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("num %d gopher, sleeping 3 seconds already\n", i)
}
