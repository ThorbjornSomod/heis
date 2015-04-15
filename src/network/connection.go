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

func ConnReceive(BroadcastPort string,client string,RecevieStruct chan NetworkInterface){//Receive messages from UDP and send to channels
	addr, _ := net.ResolveUDPAddr("udp",":" + BroadcastPort)
	conn, _ := net.ListenUDP("udp4", addr)
	if client == "slave"{
		
		for{
			conn.SetReadDeadline(time.Now().Add(1*time.Second))
			b := make([]byte,1024)
			_, _ , err := conn.ReadFromUDP(b)
			if err == nil{}	
		}	
	}
	if client == "master"{
		for{
			b := make([]byte,1024)
			length, _ , err := conn.ReadFromUDP(b)
			b = b[0:length]
			if err == nil{ 
				var m NetworkInterface
				json.Unmarshal(b,&m)
				ReceiveStruct <- m
				 
			}
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

func CostFunction(){
		
}

func Master(ReceiveStruct chan NetworkInterface){
	for{
		tempStruct :=<- ReceiveStruct
		tempIPlist := tempStruct.Message
		Println("yolo")
		Println(tempIPlist)
	}
	/*	- lager execution list og sender til network unit
		- ser om en slave er død
		- lage IP list
	*/
}

func test(IPchan chan string){
	for{
	IPchan <- "1"
	}
}
func test2(IPlistChan chan []string){
	for{
	a := <-IPlistChan
	Println(a)
	}
}
func test3(ExecuteListChan chan []int){
	array := [4]int{2,1,3,1}
	ExecuteListChan <- array[0:] 
}





func Network(){
	BroadcastIP, BroadcastPort,MyIP,client := Init()
	go test(IPchan)
	go test2(IPlistChan)
	go test3(ExecuteListChan)

	go CreateStruct(InternalOrdersToNetwork,ExternalOrdersToNetwork,MyIP,StructChannel,Direction,FloorChan)
	go Master(ReceiveStruct)			

	go ConnReceive(BroadcastPort,client,ReceiveStruct)
	go ConnSend(BroadcastPort,BroadcastIP,StructChannel)
	go MakeIPList(IPlistChan, IPchan, MyIP)

	time.Sleep(100*time.Millisecond)


			
	
	channela := make(chan string)		
	<- channela

}