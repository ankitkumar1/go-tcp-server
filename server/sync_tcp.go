package server

import (
	"errors"
	"go-tcp-server/config"
	"io"
	"log"
	"net"
	"strconv"
)

func RunTcpSyncServer() {
	address := config.Host + ":" + strconv.Itoa(config.Port)
	log.Println("Starting a synchronous TCP Server on", address)
	var connectedClients int = 0

	// Starting a listener over tcp which will listen on given address (Host and Port)
	lstnr, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer lstnr.Close() // Good practice to ensure the listener closes eventually

	for {
		c, err := lstnr.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		connectedClients += 1
		log.Println("Client connected with address:", c.RemoteAddr(), "| Concurrent Connections:", connectedClients)

		// Inner loop handles the persistent client connection
		for {
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				connectedClients -= 1
				log.Println("Client disconnected:", c.RemoteAddr(), "| Concurrent Connections:", connectedClients)
				if errors.Is(err, io.EOF) {
					break
				}
				log.Println("Read error:", err)
				break
			}

			log.Println("Command received:", cmd)
			if err = respond(cmd, c); err != nil {
				log.Println("Error writing to client:", err)
			}
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