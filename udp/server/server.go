package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Start a UDP server on port 8081
	// ListenPacket doesn't require a connection to be established beforehand
	ln, err := net.ListenPacket("udp", "localhost:8081")
	if err != nil {
		fmt.Printf("Error starting UDP server: %v", err)
		return
	}
	defer ln.Close()

	fmt.Println("UDP Server is listening on port 8081...")

	// Handle incoming packets
	for {
		buffer := make([]byte, 1024)
		// ReadFrom reads packets from any cient w/o expecting a connection
		n, addr, err := ln.ReadFrom(buffer)
		if err != nil {
			fmt.Printf("Error reading data from client: %v", err)
			continue
		}

		// Simulate network by pausing the server for 10 milliseconds after receiving data
		time.Sleep(time.Millisecond * 10)

		// Process the incoming data
		clientMessage := string(buffer[:n])

		// Send the same message back to the client
		// WriteTo specifies the address to send res to the correct client
		_, err = ln.WriteTo([]byte(clientMessage), addr)
		if err != nil {
			fmt.Printf("Error sending response: %v\n", err)
			continue
		}
	}
}
