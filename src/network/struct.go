package network

import (
	//."fmt"
	"time"
)


/*
func RandSeq(n int) string{ //Function generating a random string of length n.

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b:= make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return(string(b))
}
*/

type NetworkInterface struct {
	IP string 
	NewExternalOrders [4][2]int
	NewInternalOrders [4]int
	//Direction int
	NextDirection int
	Floor int		
}

var StructChannel = make(chan NetworkInterface)
var StructListChan = make(chan [N_ELEVATORS]NetworkInterface)

func CreateStruct(InternalOrdersToNetwork chan [4]int,ExternalOrdersToNetwork chan[4][2]int, MyIP string,StructChannel chan NetworkInterface, FloorChan chan int) {
	for{
		Internal :=<- InternalOrdersToNetwork
		External :=<- ExternalOrdersToNetwork
		//dirn :=<- Direction
		floor :=<- FloorChan
		Struct := NetworkInterface{IP:MyIP, NewExternalOrders:External, NewInternalOrders:Internal, Floor:floor} 
		StructChannel <- Struct
		time.Sleep(50*time.Millisecond)
	}	
}