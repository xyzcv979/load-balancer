package server

import (
	"fmt"
	"mymodule/src"
	"mymodule/src/client"
	"net"
	"os"
	"strconv"
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

func writeToLoadBalancer(conn net.Conn) {
	_, err := conn.Write([]byte(src.ServerMsg))
	if err != nil {
		fmt.Println("Error server write:", err.Error())
	}
}

func ReadFromClient(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	msgLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error server reading:", err.Error())
	}
	clientMsg := string(buffer[:msgLen])
	//fmt.Println("Server received:", clientMsg)
	return clientMsg, err
}

func ProcessClient(conn net.Conn) {
	clientMsg, _ := ReadFromClient(conn)
	// write to back-end server
	fmt.Println(clientMsg)
	// get server response
	msg := "Server received your msg!"
	writeToClient(conn, msg)
	defer conn.Close()
}

func connectToRandomPort() {

}

func registerWithLoadBalancer() net.Conn {
	loadBalancerAddr := src.ServerHost + ":" + strconv.Itoa(src.ServerPort)
	ret, _ := client.ConnectToServer(src.ServerType, loadBalancerAddr)
	return ret
}

func sendHeartbeat() {

}

// Server connects to load balancer as a client
// Then acts as a server for load balancer to communicate with
func RunServer() {
	//server := CreateServer(src.ServerType, src.ServerRandomPort)
	//defer server.Close()
	//fmt.Println("Listening on " + server.Addr().String())
	conn := registerWithLoadBalancer()
	client.WriteToServer(conn, src.ServerMsg)
	defer conn.Close()
	for {
		fmt.Println("Server waiting for msg")
		_, err := ReadFromClient(conn)
		if err != nil {
			break
		}
	}

	//for {
	//	conn, err := server.Accept()
	//	if err != nil {
	//		fmt.Println("Error server running:", err.Error())
	//	}
	//	go ProcessClient(conn)
	//}
}
