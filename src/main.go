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
	BroadcastIP, BroadcastPort,WelcomePort,MyIP,client := Init()
	Println(BroadcastPort,BroadcastIP,WelcomePort,MyIP,client)
	Println(IPlist)
	switch  {
		case client == "master":
			Println("bbbb")
			go SendWelcomeMessage(BroadcastIP,WelcomePort)
			go AddNewClient(BroadcastIP,BroadcastPort)
		Println("a")
		for{
			time.Sleep(100*time.Millisecond)
		}
	case client == "slave":
		Println("ccccccc")
	}

	

}
