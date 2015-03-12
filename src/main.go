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
	client := "init"
	BroadcastIP, BroadcastPort,MyIP,client := Init()
	Println(MyIP)

	switch  {				
			case client == "master":
				for{
					Println(BroadcastPort)
					Println(client)
					Println(BroadcastIP)	
					go ConnReceive(BroadcastPort,client,MasterIsAlive)
					go ConnSend(BroadcastPort,BroadcastIP)
					go ImAlive(client,MasterAliveMessage)

			
					time.Sleep(100*time.Millisecond)
				}
			case client == "slave":
				for{
					go MasterAlive(MasterIsAlive)
					time.Sleep(100*time.Millisecond)
				}
			
	}		
	

}
