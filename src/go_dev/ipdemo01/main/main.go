package main

import (
	"fmt"
	"net"
)

func host(ip, cidr string) bool {
	_, ipnet, err := net.ParseCIDR(cidr)
	ipadd := net.ParseIP(ip)
	if err != nil {
		fmt.Println("cidr is error", err)
		return false
	}
	return ipnet.Contains(ipadd)
}

func main() {
	fmt.Println(host("192.168.1.1", "192.168.1.0/24"))
}
