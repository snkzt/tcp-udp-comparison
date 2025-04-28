# HTTP Project

This folder contains a simple HTTP server and client built on top of TCP.

## How to Run
1. Run `server.go` first.
2. Then run `client.go`.

## What to Observe
- HTTP uses TCP underneath, inheriting TCP’s reliability and ordering.
- HTTP adds its own structure (methods like GET, POST, headers, body) on top of TCP.
- HTTP communication has more overhead than raw TCP due to extra metadata (headers, status codes).

## Notes
HTTP is an application-layer protocol designed for communication between web clients and servers.  
It relies on TCP to handle packet delivery and order.
