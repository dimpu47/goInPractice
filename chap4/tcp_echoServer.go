package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Failed to open port 1026.")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	reader := bufio.NewReader(conn)

	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read connection from socket.")
		conn.Close()
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	defer func() {
		conn.Close()
	}()
	conn.Write(data)
}
