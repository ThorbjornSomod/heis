package network

import(
		."fmt"
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

func CostFunction(ReceiveStruct chan NetworkInterface, IPlistChan chan [N_ELEVATORS]string, StructListChan chan [N_ELEVATORS]NetworkInterface, MyIP string,ExecuteListChan chan []int){
	var internalOrders [N_ELEVATORS][4]int
	var externalOrders [4][2]int
	var ExecuteList []int
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
			dir := MyStruct.Direction
			internal := internalOrders[i]
			for j := 0; j < len(internal); j++{
				if len(ExecuteList) > 0{
					if internal[j] == 1 && contains(ExecuteList,j) == false{
						ExecuteList = append(ExecuteList[0:],j)
					}
					if internal[j] == 0 && contains(ExecuteList,j) == true{
						Println("hei")
						position := containsPosition(ExecuteList,j)
						Println(position)
						temp := ExecuteList[len(ExecuteList)-1:][0]
						Println(temp)
						ExecuteList = append(ExecuteList[:position], ExecuteList[position+1:]...)
						//ExecuteList = append(ExecuteList[0:],temp)
					}
				}else{
					if internal[j] == 1 && contains(ExecuteList,j) == false{
						ExecuteList = append(ExecuteList[0:],j)
					}
				}
			}
			/*
			for j := 0; j < len(ExecuteList); j++{
				if dir == 0{

				}
				if dir == 1{

				}
				if dir == -1{}
			}
			*/
			Println(MyStruct)
			Println(floor)
			Println(dir)
			Println(ExecuteList)
		}
	}	


		Println(externalOrders)
		time.Sleep(10*time.Millisecond)
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

func DistributeOrders(ReceiveStruct chan NetworkInterface, IPchan chan string, ExecuteListChan chan []int, IPlistChan chan [N_ELEVATORS]string, MyIP string){
	//go MakeIPList(IPchan, IPlistChan)
	go MakeLists(IPchan, IPlistChan, ReceiveStruct,StructListChan)
	go CostFunction(ReceiveStruct,IPlistChan,StructListChan,MyIP,ExecuteListChan)
}
