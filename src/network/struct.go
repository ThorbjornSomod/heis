package network

import (
	//."fmt"
	//"time"
)




type NetworkInterface struct {
	IP string 
	NewExternalOrders [4][2]int
	NewInternalOrders [4]int
	LastStop int
	NextDirection int
	Floor int
	Executed bool		
}

var StructChannel = make(chan NetworkInterface)
var StructListChan = make(chan [N_ELEVATORS]NetworkInterface)

func CreateStruct(InternalOrdersToNetwork chan [4]int,ExternalOrdersToNetwork chan[4][2]int, MyIP string,StructChannel chan NetworkInterface, FloorChan chan int, LastStopChannel chan int,ExecutedChannel chan bool) {
	lastStop := 0
	executed := false
	tempFloor := 0

	for{
		select{
			case ReceiveLastStop :=<- LastStopChannel:
				lastStop = ReceiveLastStop
			case temp :=<- ExecutedChannel:
				executed = temp
			default:

				Internal :=<- InternalOrdersToNetwork
				External :=<- ExternalOrdersToNetwork
				floor :=<- FloorChan
				if floor != tempFloor{
					executed = false
				}

				Struct := NetworkInterface{IP:MyIP, NewExternalOrders:External, NewInternalOrders:Internal, Floor:floor, LastStop:lastStop, Executed:executed} 
				StructChannel <- Struct
				tempFloor = floor

		}
	}	
}