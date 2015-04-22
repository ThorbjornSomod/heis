package network

import(
		//."fmt"
		"time"
)

const N_ELEVATORS int = 3
/*
func MakeIPList( IPchan chan string,IPlistChan chan [N_ELEVATORS]string){
	var IPlist [N_ELEVATORS]string
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

		IP := <- IPchan
		Println(IP)
		for i := 0; i < len(IPlist); i++{ //Increase timer every time elevator sends a struct
			if IPlist[i] == IP{
				allreadyadded = true
				IPtimer[i] = time.Now().Add(time.Second*2).UnixNano()/int64(time.Millisecond)
			}
		}

		
		for i := 0; i < len(IPtimer); i++{ //Removes IP from list if connection is lost
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
					IPtimer[i] = time.Now().Add(time.Second*2).UnixNano()/int64(time.Millisecond)
					break
				}
			}
		}
		
		IPlistChan <- IPlist
	
	}
}
*/
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
		
		for i := 0; i < len(IPlist); i++{ //Increase timer every time elevator sends a struct
			if IPlist[i] == IP{
				allreadyadded = true
				StructList[i] = Struct
				IPtimer[i] = time.Now().Add(time.Second*2).UnixNano()/int64(time.Millisecond)
			}
		}

		
		for i := 0; i < len(IPtimer); i++{ //Removes IP from list if connection is lost
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
	
	}
}

func CostFunction(IPlistChan chan [N_ELEVATORS]string, StructListChan chan [N_ELEVATORS]NetworkInterface, MyIP string,ExecuteListChan chan []int,DirectionChan chan int){
	var internalOrders [N_ELEVATORS][4]int
	var externalOrders [4][2]int
	nextDirection := 0
	for{
		Structlist :=<- StructListChan
		IPlist :=<- IPlistChan

		for i := 0; i < len(IPlist); i++{ //Make a list of all internal and external orders in the system
			if IPlist[i] != "nil"{
				for j := 0; j < 4; j++{
				internalOrders[i][j] = Structlist[i].NewInternalOrders[j]
					for k := 0; k < 2; k++{
						externalOrders[j][k] = Structlist[i].NewExternalOrders[j][k]
					}
				}

			}
		}
		for i:= 0; i<len(IPlist);i++{
			if IPlist[i] == MyIP{
				MyStruct := Structlist[i]
				floor := MyStruct.Floor
				internal := MyStruct.NewInternalOrders

				nextDirection = 0
				closestUp := 100
				closestDown := 100
				closest := 100
				for j:= 0;j<len(internal); j++{
					if internal[j] == 1{
						if floor > j{
							if (floor-j < closestDown){
								closestDown = j
								closest = closestDown
							}
						}
						if floor < j{
							if (j-floor < closestUp){
								closestUp = j
								closest = closestUp
							}
						}	
					}
				}
				if closestDown > -1 && closestUp < 4{
					if closestUp-floor < floor-closestDown{
						closest = closestUp
					}
					if floor-closestDown < closestUp-floor{
						closest = closestDown
					}
				}
				if closestUp > 4{
					closest = closestDown
				}
				if closestDown > 4{
					closest = closestUp
				}
				if closest < 4 && closest > -1{
					if closest < floor{
						nextDirection = -1
					}
					if closest > floor{
						nextDirection = 1
					}
				}

				if closestUp-floor == floor-closestDown{
					//endre denne
					closest = closestUp
				}					
				DirectionChan <- nextDirection
				ExecuteListChan <- internal[0:]				
			}
		}

				/*
				for j := 0; j < len(internal); j++{
					if len(UnsortedExecuteList) > 0{
						if internal[j] == 1 && contains(UnsortedExecuteList,j) == false{
							UnsortedExecuteList = append(UnsortedExecuteList[0:],j)
						}
						if internal[j] == 0 && contains(UnsortedExecuteList,j) == true{
							position := containsPosition(UnsortedExecuteList,j)

							UnsortedExecuteList = append(UnsortedExecuteList[:position], UnsortedExecuteList[position+1:]...)
						}
					}else{
						if internal[j] == 1 && contains(UnsortedExecuteList,j) == false{
							UnsortedExecuteList = append(UnsortedExecuteList[0:],j)
						}
					}
				}

				/*
				closestUp := 100
				closestDown := 100
				closest := 100
				if nextDirection == 0{
					for j := 0; j < len(UnsortedExecuteList); j++{
						if UnsortedExecuteList[j] > floor{
							if (UnsortedExecuteList[j] - floor) < closestUp{
								closestUp = UnsortedExecuteList[j]
								closest = closestUp
								Println(closest)
							} 
						}
						if UnsortedExecuteList[j] < floor{
							if (floor - UnsortedExecuteList[j]) < closestDown{
								closestDown = UnsortedExecuteList[j]
								closest = closestDown
							}
						}
						if closestDown == closestUp{
							closest = closestUp
						}
					}
				}
				for j := 0; j < len(UnsortedExecuteList); j++{
					if closest < 4 && closest > -1{
						UnsortedExecuteList[0] = closest
					}

				}

				if closest < 4 && closest > -1{
					array := [1]int{closest}
					ExecuteListChan <- array[0:]
				}
				Println("closest")	
				Println(closest)
				Println(UnsortedExecuteList)
				Println(floor)
				Println("Next Direction")
				Println(nextDirection)
				*/
		
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

//arr = append(arr[1:], newElement)

func DistributeOrders(ReceiveStruct chan NetworkInterface, IPchan chan string, ExecuteListChan chan []int, IPlistChan chan [N_ELEVATORS]string, MyIP string,DirectionChan chan int){
	//go MakeIPList(IPchan, IPlistChan)
	go MakeLists(IPchan, IPlistChan, ReceiveStruct,StructListChan)
	go CostFunction(IPlistChan,StructListChan,MyIP,ExecuteListChan,DirectionChan)
}
