package main 

import (
		//"net"
		."fmt"
		"runtime"
		//"time"
		"./network"
)

const receivePort = "30000"
const sendPort = "20008"
//const myIP = GetMyIP()
//const BroadcastIP = GetBroadcastIP(myIP)

func main(){

	runtime.GOMAXPROCS(runtime.NumCPU())
	Println("Hei")
	Println(myIP)
}
