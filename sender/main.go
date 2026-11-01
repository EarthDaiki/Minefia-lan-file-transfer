package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Open a file you send
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("File Error: ", err)
		return
	}
	defer file.Close()

	// Connect to Receiver
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Connection Error: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected!")

	// Copy a file to TCP
	n, err := io.Copy(conn, file)
	if err != nil {
		fmt.Println("Send Error: ", err)
		return
	}

	fmt.Printf("Send %d bytes\n", n)
}
