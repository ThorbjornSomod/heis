package main 

import (
		//"net"
		."fmt"
		"runtime"
		//"time"
		."./network"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	IPlist := make(chan []string)
	BroadcastIP, BroadcastPort,ReceivePort,MyIP,client := Init()
	go AddNewClient(BroadcastIP,BroadcastPort,client,IPlist)
	Println("Hei")
	Println(MyIP)
	Println(BroadcastIP)
	Println(BroadcastPort)
	Println(ReceivePort)
	select{}
}
