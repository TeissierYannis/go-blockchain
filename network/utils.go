package network

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func GetPublicIP() string {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return string(ip)
}

func GetPrivateIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
