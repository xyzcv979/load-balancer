/*
Client socket
1. Connect to loadbalancer via loadbalancer host and port by dialing
2. Write to loadbalancer
3. Read from loadbalancer
*/

package client

import (
	"fmt"
	"net"
)

func WriteToServer(conn net.Conn, msg string) {
	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Failed to write to server", err.Error())
	}
}

func ReadFromServer(conn net.Conn) string {
	buffer := make([]byte, 1024)
	msgLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error client reading", err.Error())
	}
	return string(buffer[:msgLen])
}

func ConnectToServer(serverType string, address string, msg string) net.Conn {
	conn, err := net.Dial(serverType, address)
	if err != nil {
		fmt.Println("Error client dial", err.Error())
	}
	return conn
}
