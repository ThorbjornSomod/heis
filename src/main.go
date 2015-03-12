package main 

import (
		//"net"
		."fmt"
		"runtime"
		"time"
		."./network"
		."./channels"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	BroadcastIP, BroadcastPort,MyIP,client := Init()
	Println(BroadcastPort,BroadcastIP,BroadcastPort,MyIP,client)
	Println(IPlist)
	switch  {
		case client == "master":
			Println("bbbb")
			go SendOrders(BroadcastIP,BroadcastPort)
			//go AddNewClient(BroadcastIP,BroadcastPort)
		Println("a")
		for{
			time.Sleep(100*time.Millisecond)
		}
		case client == "slave":
			Println("ccccccc")
			//ReadMessage(BroadcastIP, WelcomePort)
		}

	

}
