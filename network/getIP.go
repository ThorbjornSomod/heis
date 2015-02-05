package getIP

import (
		"net"
		//."fmt"
		."strings"
		//"runtime"
		//"time"
)

func GetMyIP() string{
	addrs, _ := net.InterfaceAddrs() //returns table of interface addrs
	//Split(addrs[1],)
	return(addrs[1])
}

func GetBroadcastIP() string {

}


