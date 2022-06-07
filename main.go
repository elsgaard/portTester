package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

const (
	CONN_TYPE = "tcp"
)

func main() {
	portPtr := flag.String("port", "80", "Listening port")
	hostPtr := flag.String("host", "localhost", "bind to host")

	flag.Parse()

	l, err := net.Listen(CONN_TYPE, *hostPtr+":"+*portPtr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + *hostPtr + ":" + *portPtr)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Println("Got connection")
	conn.Close()
}
