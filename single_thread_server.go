package main

import (
	"flag"
	"go-tcp-server/config"
	"go-tcp-server/server"
	"log"
)


func main() {
	// Testing sync server
	setupFlag()
	log.Println("Starting the server!")
	server.RunTcpSyncServer()
}

func setupFlag() {
	// Passing the memory addresses of our global config variables
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the sync tcp server!")
	flag.IntVar(&config.Port, "port", 7379, "Port of the sync TCP server!")
	flag.Parse()
}
