package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, error := net.Listen("tcp", ":8080")
	if error != nil {
		log.Panic(error)
	}
	defer li.Close()

	for {
		conn, error := li.Accept()
		if error != nil {
			log.Println(error)
		}
		io.WriteString(conn, "Hello from TCP server\n")
		fmt.Fprintf(conn, "How are you?\n")
		fmt.Fprintf(conn, "%v", "well i hope...\n")
		conn.Close()
	}
}
