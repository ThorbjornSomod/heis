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
)

func ConnReceive(BroadcastPort string,client string,MasterIsAlive chan string){//Receive messages from UDP and send to channels
	addr, _ := net.ResolveUDPAddr("udp",":" + BroadcastPort)
	conn, _ := net.ListenUDP("udp4", addr)
	if client == "slave"{
		
		for{
			conn.SetReadDeadline(time.Now().Add(1*time.Second))
			b := make([]byte,1024)
			length, _ , err := conn.ReadFromUDP(b)
			b = b[0:length]
			if err == nil{
				if string(b) == "Welcome to the elevator system.\000"{
					MasterIsAlive <- string(b)
				}
			}
			if err != nil{
				MasterIsAlive <- "I am dead."
			}
		}	
	}
	if client == "master"{
		for{
			b := make([]byte,1024)
			length, _ , err := conn.ReadFromUDP(b)
			b = b[0:length]
			if err == nil{
					IPchan <- string(b)
			}
		}		
	}		
}

func ConnSend(BroadcastPort string, BroadcastIP string){
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + BroadcastPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for {
		message :=<- AliveMessage
		conn.Write([]byte(message+"\000"))


	}
}

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

func Network(){
	BroadcastIP, BroadcastPort,MyIP,client := Init()
	go test(IPchan)
	go test2(IPlistChan)
	switch{				
			case client == "master":
				go ConnReceive(BroadcastPort,client,MasterIsAlive)
				go ConnSend(BroadcastPort,BroadcastIP)
				go ImAlive(client,AliveMessage,MyIP)
				go MakeIPList(IPlistChan, IPchan, MyIP)
				time.Sleep(100*time.Millisecond)

			case client == "slave":
				go ImAlive(client,AliveMessage,MyIP)
				go ConnReceive(BroadcastPort,client,MasterIsAlive)
				go MasterAlive(MasterIsAlive)
				time.Sleep(100*time.Millisecond)

			
	}
	channela := make(chan string)		
	<- channela

}