package network

import (
	//."fmt"
	"time"
)




type NetworkInterface struct {
	IP string 
	NewExternalOrders [4][2]int
	NewInternalOrders [4]int
	LastStop int
	NextDirection int
	Floor int
	UpdatedGlobalExternalOrders [4][2]int		
}

var StructChannel = make(chan NetworkInterface)
var StructListChan = make(chan [N_ELEVATORS]NetworkInterface)

func CreateStruct(InternalOrdersToNetwork chan [4]int,ExternalOrdersToNetwork chan[4][2]int, MyIP string,StructChannel chan NetworkInterface, FloorChan chan int, LastStopChannel chan int,UpdatedGlobalExternalOrdersChannel chan [4][2]int) {
	lastStop := 0
	for{
		select{
			case ReceiveLastStop :=<- LastStopChannel:
				lastStop = ReceiveLastStop
			default:
				Internal :=<- InternalOrdersToNetwork
				External :=<- ExternalOrdersToNetwork
				updatedGlobalExternalOrders :=<- UpdatedGlobalExternalOrdersChannel
				floor :=<- FloorChan
				Struct := NetworkInterface{IP:MyIP, NewExternalOrders:External, NewInternalOrders:Internal, Floor:floor, LastStop:lastStop, UpdatedGlobalExternalOrders:updatedGlobalExternalOrders} 
				StructChannel <- Struct
		}
		time.Sleep(25*time.Millisecond) 
	}	
}