package main 

import (
		"net"
		."fmt"
		"runtime"
		"time"
)

const workStation = "8"
const myIp = "129.241.187.161"
const ipAddress = "129.241.187.136"
const ipBroadcast = "129.241.187.255"
const receivePort = "33546"
const sendPort = "20008" 



func tcpSend(ipAddress string, receivePort string,myIp string,conn *net.TCPConn){


	msg2 := "YOLO" + myIp + ":" + sendPort + "\x00"
	for{
		conn.Write([]byte(msg2))
		//Printf("HEi")
		time.Sleep(1000*time.Millisecond)
	}
}

func tcpReceive(ipAdress string, receivePort string,conn *net.TCPConn){
	addr, _ := net.ResolveTCPAddr("tcp", ipAddress + ":" + receivePort)
	listener, _ := net.ListenTCP("tcp", addr)
	listener.Accept()

	
	for{
		buf := make([]byte,1024)
		conn.Read(buf)
		Printf(string(buf))

	}



}
func main(){

	runtime.GOMAXPROCS(runtime.NumCPU())

	addr, _ := net.ResolveTCPAddr("tcp", ipAddress + ":" + receivePort)
	conn, _ := net.DialTCP("tcp", nil, addr)
	msg := "Connect to:" + myIp + ":" + receivePort + "\x00"
	conn.Write([]byte(msg))

	go tcpReceive(ipAddress, receivePort,conn)
	tcpSend(ipAddress,receivePort,myIp,conn)
	
	Printf("Done")
}