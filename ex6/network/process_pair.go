package network

import (
		"net"
		//."fmt"
		"time"
		"strings"
)

func GetMyIP() string{ //OK
	addrs, _ := net.InterfaceAddrs() //returns table of interface addrs
	return strings.Split(addrs[1].String(),"/")[0]
}

func GetBroadcastIP(MyIP string) string{  //OK
	myIP := strings.Split(MyIP,".")
	return myIP[0]+"."+myIP[1]+"."+myIP[2]+".161"
}

func MasterOrSlave(WelcomePort string) string{ //OK
	//Listen to broadcast for three seconds to see if somemone else is master
	addr, _ := net.ResolveUDPAddr("udp",":" + WelcomePort)
	conn, _ := net.ListenUDP("udp4", addr)
	_ = conn.SetReadDeadline(time.Now().Add(3*time.Second))
	client := ""
	b := make([]byte,1024)
	_,_,err := conn.ReadFromUDP(b)
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

func Init() (string,string,string){ //OK
	WelcomePort:= "30001"
	MyIP := GetMyIP()
	BroadcastIP := GetBroadcastIP(MyIP)
	client := MasterOrSlave(WelcomePort)
	return BroadcastIP, WelcomePort, client
}

func ReceivePing(WelcomePort string, IPchannel chan string,errorChan chan bool){ //Tror OK
	addr, _ := net.ResolveUDPAddr("udp",":" + WelcomePort)
	conn, _ := net.ListenUDP("udp4", addr)
	conn.SetReadDeadline(time.Now().Add(1000*time.Millisecond))
	for{
		b := make([]byte,1024)
		_, _ , err := conn.ReadFromUDP(b)
		if err == nil{
			IPchannel <- string(b)
			errorChan <- false
		}else{
			errorChan <- true
				}
	}
}

func AddNewClient(BroadcastIP string,WelcomePort string, client string,IPchannel chan string,errorChan chan bool){
	message := BroadcastIP
	if client == "master" { //Send IP address, and get confirmation that you are added to masters list
		SendMyPing(message,BroadcastIP,WelcomePort)
	} else if client == "slave" {
		ReceivePing(WelcomePort,IPchannel,errorChan)
	}		
}

func SendMyPing(message string,BroadcastIP string,WelcomePort string){ //OK
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + WelcomePort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for {
		conn.Write([]byte(message+"\000"))
		time.Sleep(1000*time.Millisecond)
	
	}
}