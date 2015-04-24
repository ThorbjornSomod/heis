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

	// signalChan := make(chan os.Signal, 1)
	// cleanupDone := make(chan bool)
	// signal.Notify(signalChan, os.Interrupt)
	// go func() {
 //    	for _ = range signalChan {
 //        	fmt.Println("\nFATAL ERROR: System rebooting.\n")
        	
 //        	// Somestruct <--SomeChannel: Receive struct over channel
 //        	// Write struct to file
 //        	// Do other cleanup
 //        	// Spawn new process
 //        	// Continue

 //        	cleanupDone <- true
 //    	}
	// }()
	// <-cleanupDone

	channela := make(chan string)		
	<- channela

}
