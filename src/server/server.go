package server

import (
	"fmt"
	"mymodule/src"
	"net"
	"os"
)

func CreateServer(serverType string, address string) net.Listener {
	server, err := net.Listen(serverType, address)
	if err != nil {
		fmt.Println("Error server", err.Error())
		os.Exit(1)
	}
	return server
}

func writeToClient(conn net.Conn, msg string) {
	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error server write:", err.Error())
	}
}

func ReadFromClient(conn net.Conn) string {
	buffer := make([]byte, 1024)
	msgLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error server reading:", err.Error())
	}
	clientMsg := string(buffer[:msgLen])
	//fmt.Println("Server received:", clientMsg)
	return clientMsg
}

func ProcessClient(conn net.Conn) {
	clientMsg := ReadFromClient(conn)
	// write to back-end server
	fmt.Println(clientMsg)
	// get server response
	msg := "Server received your msg!"
	writeToClient(conn, msg)
	defer conn.Close()
}

// Server connects to load balancer as a client
// Then acts as a server for load balancer to communicate with
func RunServer(address string) {
	server := CreateServer(src.ServerType, address)
	defer server.Close()
	fmt.Println("Listening on " + address)
	fmt.Println("Waiting for client...")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error server running:", err.Error())
		}
		go ProcessClient(conn)
	}
}
