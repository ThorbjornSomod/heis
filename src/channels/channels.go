package channels 

import (
		//"net"
		//."fmt"
		//"runtime"
		//"time"
		
)

var MasterIsAlive = make(chan string)
var AliveMessage = make(chan string)
var IPlistChan = make(chan []string)
var IPchan = make(chan string)