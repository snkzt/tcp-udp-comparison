package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var droppedMessagesFromClient int
	var droppedMessagesFromServer int
	var errorResponses int
	var totalRoundTripTime time.Duration
	const totalMessages = 10000

	// Resolve the server address
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error resolvingn address: %v\n", err)
		return
	}

	// Set up a local UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}

	defer conn.Close()

	// Send and receive messages
	for i := 0; i < totalMessages; i++ {
		// Generate a unique message ID
		messageID := fmt.Sprintf("MessageID-%d", i)

		// Record the send time
		start := time.Now()

		// Send the message with ID
		_, err := conn.Write([]byte(messageID))
		if err != nil {
			fmt.Printf("Error sending message: %v", err)
			droppedMessagesFromClient++
			continue
		}

		// Set a read timeout (optional, in case server doesn't respond)
		_ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))

		// Read the server's response
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading response: %v", err)
			errorResponses++
			continue
		}

		// Check if the server's response matches the sent message ID
		response := string(buffer[:n])
		if response != messageID {
			fmt.Printf("Mismatch in response: expected %s, got %s", messageID, response)
			droppedMessagesFromServer++
		}

		// Log the total round-trip time
		duration := time.Since(start)
		if i < 10 {
			fmt.Printf("Message %d round-trip time: %s\n", i, duration)
		}

		// Accumulate round-trip time for average calculation
		totalRoundTripTime += duration
	}

	// After all messages, log the summary
	averageRoundTripTime := totalRoundTripTime / time.Duration(totalMessages)
	fmt.Printf("\nTotal messages sent: %d\n", totalMessages)
	fmt.Printf("Total dropped messages from client: %d\n", droppedMessagesFromClient)
	fmt.Printf("Total dropped messages from server: %d\n", droppedMessagesFromServer)
	fmt.Printf("Total error response from server: %d\n", errorResponses)
	fmt.Printf("Average round-trip time: %s\n", averageRoundTripTime)
}
