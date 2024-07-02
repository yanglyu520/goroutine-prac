package main

import "fmt"

func main() {
	s := "aaa"

	sl := []byte(s)
	sl[0] = 't'

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println(s)
	fmt.Println(string(sl))
}
