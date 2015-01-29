package main 

import (
		"net"
		."fmt"
		"runtime"
		"time"
)

const workStation = "8"
const ipAddress = "129.241.187.161"
const ipBroadcast = "129.241.187.255"
const receivePort = "30000"
const sendPort = "20008" 


func udpSend(ipBroadcast string, sendPort string){
	addr, _ := net.ResolveUDPAddr("udp", ipBroadcast + ":" + sendPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for {
		conn.Write([]byte("YOLO \000"))

		time.Sleep(1000*time.Millisecond)
	}
}

func udpListen(ipBroadcast string, sendPort string){
	addr, _ := net.ResolveUDPAddr("udp", ipBroadcast + ":" + sendPort)
	conn, _ := net.ListenUDP("udp4", addr)
	b := make([]byte,1024)
	for {
		_, _ , err := conn.ReadFromUDP(b)
		if err == nil{
		Printf(string(b))
		}

	}
}

func main(){

	runtime.GOMAXPROCS(runtime.NumCPU())

	go udpSend(ipBroadcast,sendPort)
	udpListen(ipBroadcast, sendPort)
	Printf("Done")
}