package utils

import (
	"log"
	"net"
	"net/url"
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

func IsValidURL(u string) (bool, error) {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return false, err
	}

	if parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		return false, nil
	}
	return true, nil
}
