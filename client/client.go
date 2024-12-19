package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	// Skip certificate verification for simplicity (not recommended for production)
	config := &tls.Config{InsecureSkipVerify: true}

	// Connect to the TLS server
	conn, err := tls.Dial("tcp", "localhost:8443", config)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Send a message to the server
	message := "Hello from Client"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("failed to write: %v", err)
	}

	// Read the server's response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("failed to read: %v", err)
	}

	fmt.Printf("Received from server: %s\n", string(buf[:n]))
}
