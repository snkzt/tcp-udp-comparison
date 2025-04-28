# UDP Project

This folder contains a simple UDP server and client.

## How to Run
1. Run `server.go` first.
2. Then run `client.go`.

## What to Observe
- UDP does not guarantee delivery.
- UDP does not guarantee the order of packets.
- UDP is faster than TCP because it has no connection setup or delivery checks.
- On localhost, packet loss and disorder are rare without simulation.

## Notes
In real-world networks, UDP can experience packet loss and out-of-order delivery.  
In this localhost example, you may not easily see these issues unless they are manually simulated.
