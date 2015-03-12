package network

import (
		"net"
		."fmt"
		"time"
		//"strings"
		.".././channels"
)

func ConnReceive(BroadcastPort string,client string,MasterIsAlive chan string){//Receive messages from UDP and send to channels
	addr, _ := net.ResolveUDPAddr("udp",":" + BroadcastPort)
	conn, _ := net.ListenUDP("udp4", addr)
	if client == "slave"{
		conn.SetReadDeadline(time.Now().Add(1*time.Second))
		for{

			b := make([]byte,1024)
			length, _ , err := conn.ReadFromUDP(b)
			b = b[0:length]
			if err == nil{
				if string(b) == "Welcome to the elevator system.\000"{
					MasterIsAlive <- string(b)
				}
			}
			if err != nil{
				MasterIsAlive <- "Master is dead."
			}
		}	
	}		
}

func ConnSend(BroadcastPort string, BroadcastIP string){
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + BroadcastPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for {
		message :=<- MasterAliveMessage

		conn.Write([]byte(message+"\000"))
		time.Sleep(1000*time.Millisecond)
	}
}

func ImAlive(client string,MasterAliveMessage chan string){
	if client == "master"{
		message := "Welcome to the elevator system."
		MasterAliveMessage <- message
	}	
}	
func MasterAlive(MasterIsAlive chan string){
	for{
		alive := <- MasterIsAlive
		if alive == "Master is dead."{
			Println("Master is dead")
		}	
	}
}

