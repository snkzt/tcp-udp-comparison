#!/bin/bash

# Kill any previous instances running on ports 8080
fuser -k 8080/tcp 2>/dev/null
fuser -k 8081/udp 2>/dev/null

# Start TCP and UDP servers in the background
echo "Starting TCP server..."
./bin/tcp_server &
TCP_SERVER_PID=$!

echo "Starting UDP server..."
./bin/udp_server &
UDP_SERVER_PID=$!

# Allow servers to start
sleep 1

echo "Running TCP client..."
./bin/tcp_client > tcp_result.txt

echo "Running UDP client..."
./bin/udp_client > udp_result.txt

# Kill background serveres
kill $TCP_SERVER_PID
kill $UDP_SERVER_PID

echo "--- TCP Result ---"
cat tcp_result.txt

echo "--- UDP Result ---"
cat udp_result.txt
