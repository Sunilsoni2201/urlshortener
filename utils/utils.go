package utils

import (
	"log"
	"net"
)

// GetOutboundIP Get outbound ip of the machine
func GetOutboundIP() net.IP {
	connection, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatalf("coud not Dial connection on DNS error:%v", err)
	}
	defer connection.Close()

	localAddr := connection.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
