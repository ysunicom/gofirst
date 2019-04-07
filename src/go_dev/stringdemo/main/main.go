package main

import (
	"fmt"
	"net"
	"strings"
)

type LocationOfIP struct {
	ipadd string
	city  string
}

func main() {
	//LoOfIP := make([]LocationOfIP, 0)
	LoOfIP := []LocationOfIP{}
	ipadd := "211.91.120.130"
	ip := net.ParseIP(ipadd)
	strArr := strings.Split("211.91.120.0/24,武汉", ",")
	LoOfIP = append(LoOfIP, LocationOfIP{strArr[0], strArr[1]})
	LoOfIP = append(LoOfIP, LocationOfIP{strArr[0], strArr[1]})
	LoOfIP = append(LoOfIP, LocationOfIP{strArr[0], strArr[1]})
	LoOfIP = append(LoOfIP, LocationOfIP{strArr[0], strArr[1]})
	LoOfIP = append(LoOfIP, LocationOfIP{strArr[0], strArr[1]})
	for _, v := range LoOfIP {
		_, ipnet, _ := net.ParseCIDR(v.ipadd)
		if ipnet.Contains(ip) {
			fmt.Println(v.city)
		}
	}
	fmt.Println(LoOfIP[0], LoOfIP[1])
}
