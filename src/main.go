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

	/*signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
    	for _ = range signalChan {
        	fmt.Println("\nReceived an interrupt, stopping services...\n")
        	cleanup(services, c)
        	cleanupDone <- true
    	}
	}()
	<-cleanupDone*/

	channela := make(chan string)		
	<- channela

}
