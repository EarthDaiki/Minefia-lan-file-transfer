package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// run tcp server
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Listen Error: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Waiting for connection...")

	// wait for sender
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept Error: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected!")

	// Create a file to receive
	file, err := os.Create("received.txt")
	if err != nil {
		fmt.Println("File Creation Error: ", err)
		return
	}
	defer file.Close()

	// Copy data from TCP to a file
	n, err := io.Copy(file, conn)
	if err != nil {
		fmt.Println("Receive Error: ", err)
		return
	}

	fmt.Printf("Received: %d bytes\n", n)
}
