package network

import (
	"math/rand"
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
}

var StructChannel = make(chan NetworkInterface)

func CreateStruct(InternalOrdersToNetwork chan [4]int,ExternalOrdersToNetwork chan[4][3]int,RandSeq string, MyIP string,StructChannel chan NetworkInterface) NetworkInterface{
	n:=10
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b:= make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	Rand := string(b)
	External :=<- ExternalOrdersToNetwork
	Internal :=<- InternalOrdersToNetwork
	StructToMaster := NetworkInterface{RandomSequence:Rand, Message:MyIP, NewExternalOrders:External, NewInternalOrders:Internal} 
	return(StructToMaster)
}