package network

import (
	"math/rand"
	//."fmt"
)



func RandSeq(n int) string{ //Function generating a random string of length n.

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b:= make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return(string(b))
}


type NetworkInterface struct {
	RandomSequence string 
	Message string 
	ExecuteList [][]string
	NewExternalOrders [4][3]int
	NewInternalOrders [4]int
	IPlist []string 
	Direction int
	Floor int		
}

var StructChannel = make(chan NetworkInterface)

func CreateStruct(InternalOrdersToNetwork chan [4]int,ExternalOrdersToNetwork chan[4][3]int, MyIP string,StructChannel chan NetworkInterface, Direction chan int, FloorChan chan int) {
	for{
		n:=100
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b:= make([]rune, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		Rand := string(b)
		Internal :=<- InternalOrdersToNetwork
		External :=<- ExternalOrdersToNetwork
		dirn :=<- Direction
		floor :=<- FloorChan
		StructToMaster := NetworkInterface{RandomSequence:Rand, Message:MyIP, NewExternalOrders:External, NewInternalOrders:Internal, Direction:dirn, Floor:floor} 
		StructChannel <- StructToMaster
	}	
}