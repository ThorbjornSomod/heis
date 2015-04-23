package network

import(
		."fmt"
		"time"
)

const N_ELEVATORS int = 3


func MakeLists( IPchan chan string,IPlistChan chan [N_ELEVATORS]string, ReceiveStruct chan NetworkInterface,StructListChan chan [N_ELEVATORS]NetworkInterface){
	var IPlist [N_ELEVATORS]string
	var StructList [N_ELEVATORS]NetworkInterface
	for i:=0;i<len(IPlist);i++{
		IPlist[i] = "nil"
	}
	var IPtimer [N_ELEVATORS]int64 //Timer 
	for i:=0;i<len(IPtimer);i++{
		IPtimer[i] = 0
	}			
	IPtimerCheck := IPtimer

	for{
		allreadyadded := false

		Struct := <- ReceiveStruct
		IP := <- IPchan
		Println(IP)
		
		for i := 0; i < len(IPlist); i++{ // Increase timer every time elevator sends a struct.
			if IPlist[i] == IP{
				allreadyadded = true
				StructList[i] = Struct
				IPtimer[i] = time.Now().Add(time.Second*2).UnixNano()/int64(time.Millisecond)
			}
		}

		
		for i := 0; i < len(IPtimer); i++{ // Removes IP from list if connection is lost.
			if IPlist[i] != "nil"{
				IPtimerCheck[i] = time.Now().UnixNano()/int64(time.Millisecond)
				if IPtimer[i] < IPtimerCheck[i] && IPlist[i] != "nil"{
					IPlist[i] = "nil"
				}
			}	
		}

		if allreadyadded == false{	
			for i:=0;i<len(IPlist); i++{
				if IPlist[i] == "nil"{
					IPlist[i] = IP
					StructList[i] = Struct
					IPtimer[i] = time.Now().Add(time.Second*2).UnixNano()/int64(time.Millisecond)
					break
				}
			}
		}
		StructListChan <- StructList
		IPlistChan <- IPlist
		Println(IPlist)
	}
}

func CostFunction(IPlistChan chan [N_ELEVATORS]string, StructListChan chan [N_ELEVATORS]NetworkInterface, MyIP string,ExecuteListChan chan []int,DirectionChan chan int){
	var internalOrders [N_ELEVATORS][4]int
	var externalOrders [4][2]int
	nextDirection := 0
	for{
		Structlist :=<- StructListChan
		IPlist :=<- IPlistChan

		for i := 0; i < len(IPlist); i++{ // Make a list of all internal and external orders in the system.
			if IPlist[i] != "nil"{
				for j := 0; j < 4; j++{
				internalOrders[i][j] = Structlist[i].NewInternalOrders[j]
					for k := 0; k < 2; k++{
						externalOrders[j][k] = Structlist[i].NewExternalOrders[j][k]
					}
				}

			}
		}

// What every elevator should do.
//---------------------------------------------------------------------------------------------------		

		for i:= 0; i<len(IPlist);i++{
			if IPlist[i] == MyIP{ // For loop only active for me.

				MyStruct := Structlist[i]
				floor := MyStruct.Floor
				internal := MyStruct.NewInternalOrders
				lastStop := MyStruct.LastStop
				closestUp := 100
				closestDown := 100
				closest := 100

// Managing cost of internal orders.
//---------------------------------------------------------------------------------------------------

				for j:= 0;j<len(internal); j++{
					if internal[j] == 1{

						if floor > j{
							if (floor-j < closestDown){ // Least cost down.
								closestDown = j
								closest = closestDown
							}
						}
						if floor < j{
							if (j-floor < closestUp){ // Least cost up.
								closestUp = j
								closest = closestUp
							}
						}	
					}
				}

// Managing cost of external orders.
//---------------------------------------------------------------------------------------------------

				for j:=0;j<4;j++{
					if externalOrders[j][0] == 1{
						if (floor-externalOrders[j][0]) < closestUp{
							closestUp = j
							internal[j] = 1
						}
					}
					if externalOrders[j][1] == 1{
						if (externalOrders[j][1])-floor < closestDown{
							closestDown = j
							internal[j] = 1
						}	
					} 
				}

//---------------------------------------------------------------------------------------------------

				if closestUp-floor == floor-closestDown{ // Prioritize the order in direction of travel.
					if lastStop < floor{
						closest = closestUp
					}else if lastStop > floor{
						closest = closestDown
					}
				}else if closestUp-floor < floor-closestDown{
					closest = closestUp
				}else if closestUp-floor > floor-closestDown{
					closest = closestDown 
				}

				if lastStop < floor{
					closest = closestUp
				}else if lastStop > floor{
					closest = closestDown
				}

				if closestUp > 4{ // If no order up.
					closest = closestDown
				}else if closestDown > 4{ // If no order down.
					closest = closestUp
				}	
				
				if closest < 4 && closest > -1{
					if closest < floor{

						nextDirection = -1
					}
					if closest > floor{
						nextDirection = 1
					}
				}else{
					nextDirection = 0
				}
				DirectionChan <- nextDirection
				ExecuteListChan <- internal[0:]
				Println("")								
			}
		}
	}
}

func contains(s []int, e int) bool {
    for _, a := range s { if a == e { return true } }
    return false
}

func containsPosition(s []int, e int) int {
    for i:= 0; i< len(s); i++{
    	if s[i] == e{
    		return i
    	}
    }
    return -1
}

func DistributeOrders(ReceiveStruct chan NetworkInterface, IPchan chan string, ExecuteListChan chan []int, IPlistChan chan [N_ELEVATORS]string, MyIP string,DirectionChan chan int){
	go MakeLists(IPchan, IPlistChan, ReceiveStruct,StructListChan)
	go CostFunction(IPlistChan,StructListChan,MyIP,ExecuteListChan,DirectionChan)
}
