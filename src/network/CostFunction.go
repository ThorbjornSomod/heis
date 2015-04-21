package network

import(
		."fmt"
		"time"
)

const N_ELEVATORS int = 3

func MakeIPList( IPchan chan string,IPlistChan chan []string){
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

		for i := 0; i < len(IPlist); i++ { //Increase timer every time elevator sends a struct
			if IPlist[i] == IP{
				allreadyadded = true
				IPtimer[i] = time.Now().Add(time.Second*2).UnixNano()/int64(time.Second)
				
			}
		}

		
		for i := 0; i < len(IPtimer); i++{ //Removes IP from list if connection is lost
			IPtimerCheck[i] = time.Now().UnixNano()/int64(time.Second)
			if IPtimer[i] < IPtimerCheck[i] && IPlist[i] != "nil"{
				IPlist[i] = "nil"
				Println("hei")
			}

		}

		if allreadyadded == false{	
			for i:=0;i<len(IPlist); i++{
				if IPlist[i] == "nil"{
					IPlist[i] = IP
					break
				}
			}
		}
		
		IPlistChan <- IPlist[0:]
	
	}
}


func CostFunction(ReceiveStruct chan NetworkInterface, IPlistChan chan []string){
	for{
		Struct :=<- ReceiveStruct
		IPlist :=<- IPlistChan
		Println(IPlist)
		Println(Struct)
		time.Sleep(10*time.Millisecond)
	}

}


func DistributeOrders(ReceiveStruct chan NetworkInterface, IPchan chan string, ExecuteListChan chan []int, IPlistChan chan []string){
	go MakeIPList(IPchan, IPlistChan)
	go CostFunction(ReceiveStruct,IPlistChan)


}
