package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// Start a TCP server on port 8080
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
		return
	}
	defer ln.Close()

	fmt.Println("TCP Server is listening on port 8080...")

	// Handle incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go hadleRequest(conn)
	}
}

func hadleRequest(conn net.Conn) {
	// Ensures conn is always closed when the function returns or exits
	defer conn.Close()

	droppedMessagesFromServer := 0

	for {
		// Read incoming data into buffer slice
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed the connection")
				break
			}
			fmt.Printf("Error reading data from client: %v", err)
			droppedMessagesFromServer++
			continue
		}

		// Simulate network by pausing the server for 10 milliseconds after receiving data
		time.Sleep(time.Millisecond * 10)

		// Process the incoming data
		clientMessage := string(buffer[:n])

		// Format the message and write back to the connection
		_, err = conn.Write([]byte(clientMessage))
		if err != nil {
			fmt.Printf("Error sending response: %v", err)
			droppedMessagesFromServer++
			continue
		}
	}
}
