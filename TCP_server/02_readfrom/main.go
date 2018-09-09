package main

import (
	"bufio"
	"fmt"
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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	defer conn.Close()
}
