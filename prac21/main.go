package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	for i := range 20 {
		go func() {
			fmt.Println(i)
			select {}
		}()
	}
	go func() {
		http.ListenAndServe("8080", nil)
	}()

	select {}
}
