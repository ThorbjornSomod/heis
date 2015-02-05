package network

import (
		"net"
		//."fmt"
		"time"
		"strings"
)

func GetMyIP() string{
	addrs, _ := net.InterfaceAddr() //returns table of interface addrs
	return strings.Split(addrs[1],String(),"/")[0]
	
}

func GetBroadcastIP(myIP string) string{
	myIP := Split(myIP,".")
	return myIP[0]+myIP[1]+myIP[2]+".255"
}

func BroadcastNewClient(broadcastIP,broadcastPort){
	addr, _ := net.ResolveUDPAddr("udp", broadcastIP + ":" + broadcastPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for {
		conn.Write([]byte("new client \000"))

		time.Sleep(1000*time.Millisecond)
	}
}

func AddNewClient(){
	//if no answer; i'm master. If answer; I'm slave
	//if master; add clients
}

func init(){
	GetMyIP()
	GetBroadcastIP()
	BroadcastNewClient()
	AddNewClient()
}
