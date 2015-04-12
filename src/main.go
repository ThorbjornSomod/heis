package main 

import (
		//"net"
		//."fmt"
		"runtime"
		//"time"
		."./network"
		//."./channels"
		."./elevator"
		//."./variables"
)


func main(){

	runtime.GOMAXPROCS(runtime.NumCPU())
	
	go Elevator()
	go Network()
	channela := make(chan string)		
	<- channela

}
