package network

import (
		"net"
		//"time"
		"strings"
)

func GetMyIP() string{ //OK
	addrs, _ := net.InterfaceAddrs() //returns table of interface addrs
	return strings.Split(addrs[1].String(),"/")[0]
	
}

func GetBroadcastIP(MyIP string) string{  //OK
	myIP := strings.Split(MyIP,".")
	return myIP[0]+"."+myIP[1]+"."+myIP[2]+".255"
}


func Init() (string,string,string){ //OK
	BroadcastPort := "30011"
	MyIP := GetMyIP()
	BroadcastIP := GetBroadcastIP(MyIP)
	return BroadcastIP, BroadcastPort,MyIP
}


