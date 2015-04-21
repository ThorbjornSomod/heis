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

/*
func MasterOrSlave(BroadcastPort string) string{ //OK
	//Listen to broadcast for three seconds to see if somemone else is master
	addr, _ := net.ResolveUDPAddr("udp",":" + BroadcastPort)
	conn, err := net.ListenUDP("udp4", addr)
	conn.SetReadDeadline(time.Now().Add(1*time.Second))
	client := ""
	b := make([]byte,1024)
	_,_,err = conn.ReadFromUDP(b)
	if err == nil{
		//Someone else is master, I am slave
		client = "slave"
	} else {
		//I am master
		client = "master"
	}
	conn.Close()
	return client
}
*/

func Init() (string,string,string){ //OK
	BroadcastPort := "30000"
	MyIP := GetMyIP()
	BroadcastIP := GetBroadcastIP(MyIP)
	return BroadcastIP, BroadcastPort,MyIP
}


