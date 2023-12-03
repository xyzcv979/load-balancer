/*
Server socket connection
1. Create loadbalancer and listen on host and port #
2. Accept connection from client
3. process connection
*/

package loadbalancer

import (
	"fmt"
	"mymodule/src"
	"mymodule/src/client"
	"mymodule/src/server"
	"net"
)

var serverHostAndEmptyPort = src.ServerHost + ":"

func writeToBackEndServer(msg string) net.Conn {
	// connect to backend server
	// write msg
	backendServer := server.CreateServer(src.ServerType, serverHostAndEmptyPort)
	addr := backendServer.Addr()
	serverConn := client.ConnectToServer(src.ServerType, addr.String(), msg)
	client.WriteToServer(serverConn, msg)
	fmt.Println("Load Balancer sent msg to server: ", msg)
	return serverConn
}

func readFromBackEndServer(serverConn net.Conn) string {
	return client.ReadFromServer(serverConn)
}

func processClientAndServer(clientConn net.Conn) {
	defer clientConn.Close()
	clientMsg := server.ReadFromClient(clientConn)
	// Load balancer acts as client to connect to server and send msg
	serverConn := writeToBackEndServer(clientMsg)
	defer serverConn.Close()
	serverRcvMsg := readFromBackEndServer(serverConn)
	fmt.Println("Load Balancer received msg from server: ", serverRcvMsg)
}

func createBackEndServers(numOfServers int) []net.Listener {
	var serversArr = make([]net.Listener, numOfServers)
	for i := 0; i < numOfServers; i++ {
		serversArr[i] = server.CreateServer(src.ServerType, serverHostAndEmptyPort)
	}
	return serversArr
}

func RunLoadBalancer(serverType string, address string) {
	serv := server.CreateServer(serverType, address)
	defer serv.Close()
	fmt.Println("Listening on " + address)
	fmt.Println("Waiting for client...")
	for {
		clientConn, err := serv.Accept()
		if err != nil {
			fmt.Println("Error loadbalancer running:", err.Error())
		}
		fmt.Println("Client connection from: ", clientConn.LocalAddr())
		go processClientAndServer(clientConn)
	}
}
