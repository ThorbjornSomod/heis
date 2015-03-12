package network

import (
		"net"
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

func MasterOrSlave(BroadcastPort string) string{ //OK
	//Listen to broadcast for three seconds to see if somemone else is master
	addr, _ := net.ResolveUDPAddr("udp",":" + BroadcastPort)
	conn, err := net.ListenUDP("udp4", addr)
	conn.SetReadDeadline(time.Now().Add(1*time.Second))
	client := ""
	b := make([]byte,1024)
	_,_,err = conn.ReadFromUDP(b)
	if err == nil{
		//Someone else is master, I am slave
		client = "slave"
	} else {
		//I am master
		client = "master"
	}
	conn.Close()
	return client
}

func Init() (string,string,string,string){ //OK
	BroadcastPort := "30000"
	MyIP := GetMyIP()
	BroadcastIP := GetBroadcastIP(MyIP)
	client := MasterOrSlave(BroadcastPort)
	return BroadcastIP, BroadcastPort,MyIP,client
}


/*
func SendMyIP(MyIP,BroadcastIP,BroadcastPort){ //OK
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + BroadcastPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for {
		conn.Write([]byte(MyIP+"\000"))
		time.Sleep(1000*time.Millisecond)
	}
}
*/
/*
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
*/

/*func ReadMessage(BroadcastIP string, WelcomePort string){
	Println("Hei")
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + WelcomePort)
	conn, _ := net.ListenUDP("udp4", addr)
	msg := make([]byte, 1024)
	conn.ReadFromUDP(msg)
	Println(msg)
}*/






/*
func IPAdded(BroadcastIP string,BroadcastPort string, IPchannel chan string, IPlistChannel chan []string){ // Tror OK
	addr, _ := net.ResolveUDPAddr("udp", BroadcastIP + ":" + BroadcastPort)
	conn, _ := net.DialUDP("udp", nil,addr)
	for{
		message := <- IPchannel2
		conn.Write([]byte(message+ "\000"))
		time.Sleep(1000*time.Millisecond)	
	}
}
*/
/*
func IPListMaker(IPchannel chan string, IPchannel2 chan string, IPlistChannel chan []string, IPlist []string){ // Tror OK
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
		}else if alreadyAdded == 1{ // Brukes til "I'm alive"
			AliveChannel <- newIP
			IPchannel2 <- newIP
		}
	}
}
*/






/* Con channel, IPlist channel og Alive channel må sendes ut i main*/
