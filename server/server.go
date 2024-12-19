package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Println("Read error:", err)
		return
	}

	fmt.Printf("Received: %s\n", string(buf[:n]))
	_, err = conn.Write([]byte("Hello from Server"))
	if err != nil {
		log.Println("Write error:", err)
	}
}

func main() {
	// Load certificates
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("failed to load certificates: %v", err)
	}

	// Configure TLS
	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	// Start TLS server
	ln, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	defer ln.Close()

	fmt.Println("Server listening on port 8443...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}

		go handleConnection(conn)
	}
}
