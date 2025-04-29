# TCP vs UDP - Comparison Projects

This repository demonstrates basic differences between TCP and UDP using small Go projects.

## Contents
- `tcp/`: Direct TCP communication showcasing reliable, ordered data transfer.
- `udp/`: Direct UDP communication showcasing faster but unreliable data transfer.
- `bin/`: Contains compiled server and client binaries for both protocols.
- `run_benchmark.sh`: A script to run both TCP and UDP projects simultaneously for comparison.

## Repo structure
```
/TCP-vs-UDP-Comparison
├── bin/
│   ├── tcp_server  (compiled TCP server binary)
│   ├── tcp_client  (compiled TCP client binary)
│   ├── udp_server  (compiled UDP server binary)
│   └── udp_client  (compiled UDP client binary)
├── tcp/
│   ├── client.go
│   └── server.go
│   └── README.md  (TCP project instructions)
├── udp/
│   ├── client.go
│   └── server.go
│   └── README.md  (UDP project instructions)
├── run_benchmark.sh  (shell script to run both TCP and UDP benchmarks)
└── README.md  (root README file)
```

## How to Use

### 1. Build TCP and UDP Projects
Make sure Go is installed. 
Then create bin directory run:

```bash
go build -o ./bin/tcp_server ./tcp/server/server.go
go build -o ./bin/tcp_client ./tcp/client/client.go
go build -o ./bin/udp_server ./udp/server/server.go
go build -o ./bin/udp_client ./udp/client/client.go
```

### 2. Run Both Projects for Comparison
Use the provided shell script:
```
chmod +x run_benchmark.sh
./run_benchmark.sh
```
This script will:
  - Kill any process using ports 8080 (TCP) or 8081 (UDP)
  - Launch TCP and UDP servers in the background
  - Run both clients
  - Print the results of each in separate sections

### Output
You will see output like:

```
--- TCP Result ---
<output from TCP client>

--- UDP Result ---
<output from UDP client>
```

### Notes
If you encounter "address already in use" errors, ensure no other processes are using ports 8080 or 8081.
