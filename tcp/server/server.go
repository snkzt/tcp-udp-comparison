package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// Start a TCP server on port 8080
	ln, err := net.Listen("tcp", "localhost:8080")
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

	// Read the data from the client
	buffer := make([]byte, 1024)
	droppedMessagesFromServer := 0

	for {
		// Read incoming data into buffer slice
		n, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading data from client: %v", err)
			droppedMessagesFromServer++
			continue
		}

		// Process the message
		clientMessage := string(buffer[:n])

		time.Sleep(time.Millisecond * 10)

		// Format the message and write back to the connection
		_, err = conn.Write([]byte(clientMessage))
		if err != nil {
			fmt.Printf("Error sending response: %v", err)
			droppedMessagesFromServer++
			continue
		}
	}
}
