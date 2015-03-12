package channels 

import (
		//"net"
		//."fmt"
		//"runtime"
		//"time"
		
)

var MasterIsAlive = make(chan string)
var MasterAliveMessage = make(chan string)
var IPlist = make(chan []string)