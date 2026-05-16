package main

import (
	"log"
	"net"
	"errors"
	"io"
)

var connectedClients int = 0

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Waiting for a client to connect")
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		connectedClients += 1

		log.Println("Client connected with address:", connection.RemoteAddr(), "| Concurrent Connections:", connectedClients)

		go process(connection)
	}
}

func process(connection net.Conn) {
	// Inner loop handles the persistent client connection
	for {
		cmd, err := readCommand(connection)
		if err != nil {
			connection.Close()
			connectedClients -= 1
			log.Println("Client disconnected:", connection.RemoteAddr(), "| Concurrent Connections:", connectedClients)
			if errors.Is(err, io.EOF) {
				break
			}
			log.Println("Read error:", err)
			break
		}

		log.Println("Command received:", cmd)
		if err = respond(cmd, connection); err != nil {
			log.Println("Error writing to client:", err)
		}
	}
}

// readCommand moved outside of RunTcpSyncServer and fixed return types (string, error)
func readCommand(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)

	n, err := c.Read(buf[:]) // n is number of bytes read
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

// respond moved outside, added necessary parameters (cmd, c) and fixed return type
func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}