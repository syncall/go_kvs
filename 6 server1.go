package main

import (
	"bufio"
	"fmt"
	"net"
	"runtime"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func handleClient(conn net.Conn) {
	buf := bufio.NewReader(conn)

	conn.Write([]byte("What's your name? "));

	for {
		name, err := buf.ReadString('\n')

		if err != nil {
			fmt.Printf("Client disconnected.\n")
			break
		}

		conn.Write([]byte("Hello, " + name))
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	ln, err := net.Listen("tcp", ":8080")
	check(err, "Server is ready: Connenct with telnet localhost 8080")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		go handleClient(conn);
	}
}