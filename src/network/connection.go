package network

import (
		"net"
		."fmt"
		"time"
		//"strings"
		.".././channels"
		//."encoding/json"
		//.".././elevator"
		//.".././variables"
		"encoding/json"
		//."strings"
)

var SendStruct = make(chan NetworkInterface)
var ReceiveStruct = make(chan NetworkInterface)
var IPlistChan = make(chan [N_ELEVATORS]string)

func ConnReceive(BroadcastPort string,RecevieStruct chan NetworkInterface){//Receive messages from UDP and send to channels
	addr, _ := net.ResolveUDPAddr("udp",":" + BroadcastPort)
	conn, _ := net.ListenUDP("udp", addr)	
	for{	
		conn.SetReadDeadline(time.Now().Add(time.Second*2))
		b := make([]byte,1024)
		length, _ , err := conn.ReadFromUDP(b)
		b = b[0:length]
		if err == nil{ 
			var m NetworkInterface
			json.Unmarshal(b,&m)
			ReceiveStruct <- m
			IPchan <- m.IP	 
		}else{
			Println("Something is wrong")
		}		
	}		
}

func ConnSend(BroadcastPort string, BroadcastIP string,NetworkChannel chan NetworkInterface){
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + BroadcastPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for {
		information := <- StructChannel				
		message,_ := json.Marshal(information)	
		conn.Write(message)
	}
}

func test3(ExecuteListChan chan []int){
	for{
		a :=<- ExecuteListChan
		Println(a)
		time.Sleep(100*time.Millisecond) 
	}
}

func Network(){
	BroadcastIP, BroadcastPort,MyIP := Init()

	go CreateStruct(InternalOrdersToNetwork,ExternalOrdersToNetwork,MyIP,StructChannel,FloorChan,LastStopChannel)			
	go ConnReceive(BroadcastPort,ReceiveStruct)
	go ConnSend(BroadcastPort,BroadcastIP,StructChannel)
	go DistributeOrders(ReceiveStruct, IPchan, ExecuteListChan, IPlistChan,MyIP,DirectionChan)


	channela := make(chan string)		
	<- channela

}