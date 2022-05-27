package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

// Test connectivity to SFTP server.
// Usage: sftptest <host>:<port> <user> <password>
// Example: sftptest ftp.example.com:22 root password
// Note: Make sure port number is specified.
func main() {
	// Retrieve host, username and password from command line arguments.
	args := os.Args
	if len(args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: sftptest <host>:<port> <user> <password>\nExample: sftptest ftp.example.com:22 root password")
		os.Exit(1)
	}

	host := args[1]
	username := args[2]
	password := args[3]

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// `InsecureIgnoreHostKey` is not recommended. Consider other more secured method for verifying host key
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Establishing connection
	conn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to dial: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connection successful")
}
