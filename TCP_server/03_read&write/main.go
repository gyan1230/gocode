package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	conn.SetDeadline(time.Now().Add(10 * time.Second))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "You say %s \n", line)
	}
	defer conn.Close()
	fmt.Println("****CODE Reach Here******")
}
