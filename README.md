# Minefia LAN File Transfer

A lightweight file transfer project built in **Go** to explore TCP networking, file I/O, and concurrent programming.

> **Status:** Early development. The current prototype has been tested locally on a single computer. LAN transfers between separate devices have not yet been verified.

## Features

- Send file contents over a TCP connection.
- Receive and save data to a local file.
- Display the number of bytes transferred.
- Run the sender and receiver independently from the terminal.

## Tech Stack

- **Language:** Go
- **Networking:** TCP (`net.Listen`, `net.Dial`)
- **File transfer:** Go standard library (`io.Copy`)
- **Planned:** Goroutines, channels, automatic device discovery, and transfer verification

## Getting Started

### Requirements

- Go installed on your computer
- Two terminal windows

### 1. Clone the repository

```bash
git clone https://github.com/EarthDaiki/Minefia-lan-file-transfer.git
cd Minefia-lan-file-transfer
```

### 2. Prepare a sample file

Create `text.txt` in the project root:

```bash
echo "Hello from Go!" > text.txt
```

The current sender expects a file named `text.txt`. If you change the filename in the source code, adjust this step accordingly.

### 3. Start the receiver

In the first terminal:

```bash
go run ./receiver
```

### 4. Start the sender

In a second terminal, from the same project root:

```bash
go run ./sender
```

The programs use a TCP connection on the local machine. If the address or port has been changed in the source code, make sure both programs use matching settings.

### 5. Verify the result

Check the received file and compare its contents with the original. The output filename is determined by the receiver implementation.

## Project Structure

```text
Minefia-lan-file-transfer/
├── receiver/
│   └── main.go
├── sender/
│   └── main.go
├── go.mod
├── .gitignore
├── LICENSE
└── README.md
```

## Roadmap

- [x] Establish a local TCP connection between sender and receiver
- [x] Transfer a small text file
- [ ] Support multiple simultaneous connections with goroutines
- [ ] Report transfer progress using channels
- [ ] Send file metadata (name and size)
- [ ] Detect interrupted or incomplete transfers
- [ ] Verify file integrity with SHA-256
- [ ] Discover devices on the local network
- [ ] Support file transfer between separate computers
- [ ] Add a graphical user interface

## Limitations

This is currently a learning prototype, not a secure file-sharing product. It does not yet provide authentication, transport encryption, or robust recovery from interrupted transfers. Use it only in a trusted local testing environment.

## License

Licensed under the [MIT License](LICENSE).
