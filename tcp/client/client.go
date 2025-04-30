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

	// Connect to the TCP server running on localhost:8080
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
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
		n, err := conn.Read(buffer)
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
	fmt.Println("--- Summary ---")
	fmt.Printf("\nTotal messages sent: %d\n", totalMessages)
	fmt.Printf("Total dropped messages from client: %d\n", droppedMessagesFromClient)
	fmt.Printf("Total dropped messages from server: %d\n", droppedMessagesFromServer)
	fmt.Printf("Total error response from server: %d\n", errorResponses)
	fmt.Printf("Average round-trip time: %s\n", averageRoundTripTime)
}
