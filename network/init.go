package network

import (
		"net"
		."fmt"
		"time"
		"strings"
)



func GetMyIP() string{ //OK
	addrs, _ := net.InterfaceAddrs() //returns table of interface addrs
	return strings.Split(addrs[1].String(),"/")[0]
	
}

func GetBroadcastIP(MyIP string) string{  //OK
	myIP := strings.Split(MyIP,".")
	return myIP[0]+"."+myIP[1]+"."+myIP[2]+".255"
}

func MasterOrSlave(WelcomePort string) string{ //OK
	//Listen to broadcast for three seconds to see if somemone else is master
	addr, _ := net.ResolveUDPAddr("udp",":" + BroadcastPort)
	conn, _ := net.ListenUDP("udp4", addr)
	_ = conn.SetReadDeadline(time.Now().Add(3*time.Second))
	client := ""
	b := make([]byte,1024)
	_,_,err := conn.ReadFromUDP(b)
	if err == nil{
		//Someone else is master, I am slave
		client = "slave"
	} else {
		//I am master
		client = "master"
	}
	return client
}

func SendMyIP(MyIP,BroadcastIP,BroadcastPort){ //OK
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + BroadcastPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for {
		conn.Write([]byte(MyIP+"\000"))
		time.Sleep(1000*time.Millisecond)
	}
}

func Aknowledge(BroadcastPort string) string{ //Tror OK
	addr, _ := net.ResolveUDPAddr("udp",":" + BroadcastPort)
	conn, _ := net.ListenUDP("udp4", addr)
	_ = conn.SetReadDeadline(time.Now().Add(3*time.Second))
	b := make([]byte,1024)
	for {
		_, _ , err := conn.ReadFromUDP(b)
		if err == nil{
			ConChannel <- 0
		} else{
			ConChannel <- 1
		}
	}	

}


func SendWelcomeMessage(BroadcastIP string, BroadcastPort string){ // OK
	Println("dfjsdøf")
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + BroadcastPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for {
		conn.Write([]byte("Welcome to the elevator system \000"))
		Println("YOLO")
		time.Sleep(1000*time.Millisecond)
	}
}

func ReceiveIP(BroadcastPort string, IPchannel chan string){ //Tror OK
	addr, _ := net.ResolveUDPAddr("udp",":" + BroadcastPort)
	conn, _ := net.ListenUDP("udp4", addr)
	count := 0
	alreadyAdded := 0
	for{
		b := make([]byte,1024)
		_, _ , err := conn.ReadFromUDP(b)
		if err == nil{
			IPchannel <- string(b)
				}
		}
}

func IPListMaker(IPchannel chan string, IPchannel2 chan string, IPlistChannel chan []string){ // Tror OK
	IPlist := []string
	newIP:= ""
	alreadyAdded := 0
	for{
		newIP <- IPchannel
		length := len(IPlist)
		for i:= 0; i < length; i++{
			if newIP == IPlist[i]{
				alreadyAdded = 1
			}
		}
		if alreadyAdded == 0{
			IPlist[length] = newIP
			IPlistChannel <- IPlist
			IPchannel2 <- newIP
		}
		else if alreadyAdded == 1{ // Brukes til "I'm alive"
			AliveChannel <- newIP
			IPchannel2 <- newIP
		}
	}
}

func IPAdded(BroadcastIP string,BroadcastPort string, IPchannel chan string, IPlistChannel chan []string){ // Tror OK
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + BroadcastPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for{
		message := <- IPchannel2
		conn.Write([]byte(message+ "\000"))
		time.Sleep(1000*time.Millisecond)	
	}
}


func Init() (string,string,string,string,string){ //OK
	BroadcastPort := "20008"
	WelcomePort:= "30001"
	MyIP := GetMyIP()
	BroadcastIP := GetBroadcastIP(MyIP)
	client := MasterOrSlave(WelcometPort)
	return BroadcastIP, BroadcastPort,ReceivePort,MyIP,client
}

func AddNewClient(BroadcastIP string,BroadcastPort string, client string,IPlistchannel chan []string,WelcomePort){
	Println(client)
	if client == "master"{ // Send message so that others know they are slave, receive IPs and send confirmation
		IPchannel := make(chan string)
		IPchannel2 := make(chan string)
		AliveChannel := make(chan string)
		Println("Hellu")	
			go ReceiveIP(BroadcastPort,IPchannel)
			go IPAdded()
			go IPlistMaker(IPchannel,IPListChannel)
			SendWelcomeMessage(BroadcastIP, WelcomePort)
	}
	else if client == "slave" { //Send IP address, and get confirmation that you are added to masters list
		ConfirmationChannel := make(chan string)
		SendMyIP(MyIP)
		go Aknowledge()
	}		
	
}
/* Con channel, IPlist channel og Alive channel må sendes ut i main*/
