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
	serverConn, _ := client.ConnectToServer(src.ServerType, addr.String())
	client.WriteToServer(serverConn, msg)
	fmt.Println("Load Balancer sent msg to server: ", msg)
	return serverConn
}

func readFromBackEndServer(serverConn net.Conn, servAddrMap *[]string) string {
	return client.ReadFromServer(serverConn)
}

func processClientAndServer(serv net.Listener, conn net.Conn, servAddrMap *[]net.Conn) {
	clientMsg, _ := server.ReadFromClient(conn)

	// Connection from backend server
	if clientMsg == src.ServerMsg {
		handleServerRegistration(servAddrMap, conn)
	} else { // client connection
		defer conn.Close()
		// process client and send msg to server

	}
	//// Load balancer acts as client to connect to server and send msg
	//serverConn := writeToBackEndServer(clientMsg)
	//defer serverConn.Close()
	//serverRcvMsg := readFromBackEndServer(serverConn)
	//fmt.Println("Load Balancer received msg from server: ", serverRcvMsg)
}

func processClient(conn net.Conn, servAddrMap *[]net.Conn, clientMsg string) {
	// load balancing algorithm to pick server from server address map
	for _, servConn := range *servAddrMap {
		_, err := servConn.Write([]byte(clientMsg))
		if err != nil {
			fmt.Println("Failed to write to server", err.Error())
		} else {
			break
		}
	}
}

func createBackEndServers(numOfServers int) []net.Listener {
	var serversArr = make([]net.Listener, numOfServers)
	for i := 0; i < numOfServers; i++ {
		serversArr[i] = server.CreateServer(src.ServerType, serverHostAndEmptyPort)
	}
	return serversArr
}

func handleServerRegistration(servAddrMap *[]net.Conn, conn net.Conn) {
	*servAddrMap = append(*servAddrMap, conn)
}

func handleHeartbeat() {

}

// Run loadbalancer
// Wait for client to connect
// Wait for servers to connect
// Send client requests to backend servers
// Sendd server responses back to client
func RunLoadBalancer(serverType string, address string) {
	serv := server.CreateServer(serverType, address)
	defer serv.Close()
	fmt.Println("Listening on " + address)
	fmt.Println("Waiting for client...")

	var servAddrMap []net.Conn
	for {
		conn, err := serv.Accept()
		if err != nil {
			fmt.Println("Error loadbalancer running:", err.Error())
		}
		fmt.Println("Connection from: ", conn.RemoteAddr())
		processClientAndServer(serv, conn, &servAddrMap)
		fmt.Println(servAddrMap)
		// servAddrMap = handleServerRegistration(servAddrMap, conn.RemoteAddr().String())

	}
}
