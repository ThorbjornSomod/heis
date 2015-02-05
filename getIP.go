package getIP

import (
		"net"
		."fmt"
		//."strings"
		//"runtime"
		//"time"
)

func GetMyIP() {
	addrs, _ := net.InterfaceAddrs() //returns table of interface addrs
	Println(addrs)
	
}

func GetBroadcastIP() string {

}


