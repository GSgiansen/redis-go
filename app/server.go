package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	const keyword = "+PONG\r\n"

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConn(c)

	}

}

func handleConn(c net.Conn) error {
	defer c.Close()

	// Read from the connection
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			return err
		}
		fmt.Printf("Received: %s\n", buf[:n])

		// Write to the connection
		_, err = c.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error writing: ", err.Error())
			return err
		}
	}
	return nil
}
