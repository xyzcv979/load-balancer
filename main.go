package main

import (
	"fmt"
	"mymodule/src"
	"mymodule/src/client"
	"mymodule/src/loadbalancer"
	"mymodule/src/server"
	"os"
	"strconv"
)

func main() {
	addr := src.ServerHost + ":" + strconv.Itoa(src.ServerPort)
	if len(os.Args) < 2 {
		fmt.Println("Not enough args, enter 'client' or 'loadbalancer'")
		os.Exit(1)
	}
	arg1 := os.Args[1]
	//if arg1 == "client" {
	//	client.ConnectToServer()
	//} else
	if arg1 == "loadbalancer" {
		loadbalancer.RunLoadBalancer(src.ServerType, addr)
	} else if arg1 == "server" {
		server.RunServer()
	} else if arg1 == "client" {
		client.RunClient()
	} else {
		fmt.Println("Invalid arg, enter 'client' or 'loadbalancer' or 'server'")
	}

}
