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
		time.Sleep(100*time.Millisecond)
	}
}

/*
func ImAlive(client string,MasterAliveMessage chan string,MyIP string){
	for{
		if client == "master"{
			message := "Welcome to the elevator system."
			AliveMessage <- message
		}
		if client == "slave"{
			message := "IP"+MyIP
			AliveMessage <- message
		}	
		time.Sleep(100*time.Millisecond)
	}	
}

func MasterAlive(MasterIsAlive chan string){
	for{
		alive := <- MasterIsAlive
		if alive == "I am dead."{
			Println("I am dead")	
		}	
	}
}
*/
/*
func MakeIPList(IPlistchan chan []string, IPchan chan string,MyIP string){
	var IPlist [1]string
	IPlist[0] = MyIP

	for {
		var temp [len(IPlist)+1]string
		allreadyadded := 0
		IP := <- IPchan
		for i := 0; i < len(IPlist); i++ {
			if IPlist[i] == IP{
				allreadyadded = 1
			}
		}
		if allreadyadded == 0{	
			
			temp[len(IPlist)] = IP
			for i:=0;i<len(IPlist); i++{
				temp[i] = IPlist[i]
			}	
			
		}
		IPlist := temp
		IPlistChan <- IPlist[0:]
		time.Sleep(100*time.Millisecond)	
	}
}
*/



func test3(ExecuteListChan chan []int){
	array := [4]int{2,1,3,1}
	ExecuteListChan <- array[0:] 
}





func Network(){
	BroadcastIP, BroadcastPort,MyIP := Init()
	go test3(ExecuteListChan)

	go CreateStruct(InternalOrdersToNetwork,ExternalOrdersToNetwork,MyIP,StructChannel,Direction,FloorChan)			
	go ConnReceive(BroadcastPort,ReceiveStruct)
	go ConnSend(BroadcastPort,BroadcastIP,StructChannel)
	go DistributeOrders(ReceiveStruct, IPchan, ExecuteListChan, IPlistChan)

	time.Sleep(100*time.Millisecond)


			
	
	channela := make(chan string)		
	<- channela

}