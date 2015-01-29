package main 

import (
		"net"
		."fmt"
		//"runtime"
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
		send, _:= conn.Write([]byte("YOLO"))
		Printf(string(send))
		time.Sleep(1000*time.Millisecond)
	}
}

func udpListen(ipBroadcast string, receivePort string){
	addr, _ := net.ResolveUDPAddr("udp", ipBroadcast + ":" + receivePort)
	conn, _ := net.ListenUDP("udp", addr)
	b := make([]byte,1024)
	for {
		n, _ , _:= conn.ReadFromUDP(b)
		Printf(string(b))
		Printf(string(n))

	}
}

func main(){
	udpSend(ipBroadcast,sendPort)
	udpListen(ipBroadcast, receivePort)
	Printf("Done")
}