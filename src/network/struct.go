package network

var StructChannel = make(chan NetworkInterface)

type NetworkInterface struct {
	RandomSequence, Message string 
	ExecuteList,NewOrders [][]string
	IPlist []string  		
}
