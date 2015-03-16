package main 

import (
		//"net"
		."fmt"
		"runtime"
		"time"
		."./network"
		."./channels"
		."./elevator"
		."./variables"
)

	func test(IPchan chan string){
		for{
		IPchan <- "1"
		}
	}
	func test2(IPlistChan chan []string){
		for{
		a := <-IPlistChan
		Println("Heis")
		Println(a)
	}
	}	

func main(){

	runtime.GOMAXPROCS(runtime.NumCPU())
	BroadcastIP, BroadcastPort,MyIP,client := Init()
	Println(MyIP)
	go test(IPchan)
	go test2(IPlistChan)
	a := RandSeq(MessageIdLength)
	Println(a)
	Println(client)
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
