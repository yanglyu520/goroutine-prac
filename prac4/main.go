package main

import (
	"io"
	"log"
	"net"
	"time"
)

// to test, run `nc localhost 8000`
// netcat a versatile networking utility used for reading from and writing to network connections using TCP or UDP protocols
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:01:01\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
