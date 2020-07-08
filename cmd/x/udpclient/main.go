package main

import (
	"fmt"
	"github.com/hellgate75/udpsample/pkg/common"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s host:port msg\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	message := os.Args[2]

	udpAddr, err := net.ResolveUDPAddr(common.LookupNetwork, service)
	common.CheckError(err)

	conn, err := net.DialUDP(common.Network, nil, udpAddr)
	common.CheckError(err)

	_, err = conn.Write([]byte(message))
	common.CheckError(err)

	var buf [1024]byte
	n, err := conn.Read(buf[0:])
	common.CheckError(err)

	defer func() {
		_ = conn.Close()
	}()

	fmt.Println("Answer:", string(buf[0:n]))

	os.Exit(0)
}