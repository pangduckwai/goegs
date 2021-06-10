package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	const url = "localhost:54321"

	if !((len(os.Args) == 2 && os.Args[1] == "server") || (len(os.Args) == 3 && os.Args[1] == "client")) {
		log.Fatalln("Usage: go run net [server | client]")
	}

	switch os.Args[1] {
	case "server":
		listener, err := net.Listen("tcp", url)
		if err != nil {
			log.Fatalf("Listen() failed: %v\n", err)
		}
		defer listener.Close()

		var status string

		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					fmt.Printf("Accept(): %v\n", err)
					break
				}

				go func(c net.Conn) {
					fmt.Fscanln(c, &status)
					c.Close()
				}(conn)
			}
		}()

		for i := 0; i < 60; i++ {
			time.Sleep(3 * time.Second)
			fmt.Printf("%2v: processing '%v'...\n", i, status)
		}

	case "client":
		conn, err := net.Dial("tcp", url)
		if err != nil {
			log.Fatalf("Dial() failed: %v\n", err)
		}
		defer conn.Close()

		fmt.Fprintln(conn, os.Args[2])
	}
}
