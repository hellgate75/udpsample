package main

import (
	"fmt"
	"github.com/hellgate75/udpsample/pkg/common"
	"net"
	"os"
	"strings"
)

var running = true

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr(common.LookupNetwork, service)
	common.CheckError(err)

	for running {
		conn, err := net.ListenUDP(common.Network, udpAddr)
		if err != nil {
			fmt.Printf("Connection error: %v\n", err)
			if conn != nil {
				_ = conn.Close()
			}
			continue
		}
		handleClient(conn)
		_ = conn.Close()
	}
}

func handleClient(conn *net.UDPConn) {

	var buf [1024]byte

	n, addr, err := conn.ReadFromUDP(buf[0:])

	if err != nil {
		fmt.Printf("Error reading data: %v\n", err)
		return
	}

	var cmd = fmt.Sprintf("%s", buf[0:n])
	fmt.Println("Request:", cmd)

	_, err = conn.WriteToUDP([]byte(fmt.Sprintf("Received: %s", cmd)), addr)

	if err != nil {
		fmt.Printf("Error writing data: %v\n", err)
		return
	}

	if strings.ToLower(cmd) == "exit" {
		running = false
		fmt.Println("Exit server on client request ....")
	}
}