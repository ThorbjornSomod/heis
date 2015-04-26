package elevator

import(
		."fmt"
		"time"
		"io/ioutil"
    	"encoding/json"
    	"os"
		//.".././network"
		.".././channels"
		)

const N_BUTTONS int = 3
const N_FLOORS int = 4

type elev_button_type_t int
const(
	BUTTON_CALL_UP = 0
	BUTTON_CALL_DOWN = 1
	BUTTON_COMMAND = 2
	)
type elev_motor_direction_t int
const(
	DIRN_DOWN = -1
	DIRN_STOP = 0
	DIRN_UP = 1
	)

var lamp_channel_matrix = [N_FLOORS][N_BUTTONS]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var button_channel_matrix = [N_FLOORS][N_BUTTONS]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

var dirn int

func elev_init() bool{ 
	if io_init() == 0{ //initialization of i/o
		return false
	}
	for i:=0;i<N_FLOORS;i++{
		if i!=0{
			elev_set_button_lamp(BUTTON_CALL_DOWN, i, false) //set all call down buttons to 0
		}
		if i!=N_FLOORS-1{
			elev_set_button_lamp(BUTTON_CALL_UP, i, false)	//set all call up buttons to 0
		}		
		elev_set_button_lamp(BUTTON_COMMAND, i, false) //set all floor buttons to 0
	}
	elev_set_stop_lamp(false)
	elev_set_door_open_lamp(false)
	elev_set_floor_indicator(0)
	if elev_get_floor_sensor_signal() != 0{
		elev_set_motor_direction(DIRN_DOWN)
		for{
			if elev_get_floor_sensor_signal() != -1{
				elev_set_motor_direction(DIRN_UP)
				time.Sleep(10*time.Millisecond)					
				elev_set_motor_direction(DIRN_STOP)
				dirn = DIRN_STOP
				break
			}
		}			
	}

	return true
}

func elev_set_button_lamp(button elev_button_type_t, floor int, value bool){
	if value == true{
		io_set_bit(lamp_channel_matrix[floor][button])
	}
	if value == false{
		io_clear_bit(lamp_channel_matrix[floor][button])
	}
}

func elev_set_stop_lamp(value bool){
	if value == true{
		io_set_bit(LIGHT_STOP)
	}else{
		io_clear_bit(LIGHT_STOP)
	}
}

func elev_set_door_open_lamp(value bool){
	if value == true{
		io_set_bit(LIGHT_DOOR_OPEN)
	}else{
		io_clear_bit(LIGHT_DOOR_OPEN)
	}	
}
func elev_set_floor_indicator(floor int){
	switch floor{
	case 0:
		io_clear_bit(LIGHT_FLOOR_IND1)
		io_clear_bit(LIGHT_FLOOR_IND2)
	case 1:
		io_clear_bit(LIGHT_FLOOR_IND1)
		io_set_bit(LIGHT_FLOOR_IND2)
	case 2:
		io_set_bit(LIGHT_FLOOR_IND1)
		io_clear_bit(LIGHT_FLOOR_IND2)
	case 3:
		io_set_bit(LIGHT_FLOOR_IND1)
		io_set_bit(LIGHT_FLOOR_IND2)			
	}
}

func elev_set_motor_direction(dirn elev_motor_direction_t){
	if dirn == 0{
		io_write_analog(MOTOR, 0)
	}else if dirn > 0{
		io_clear_bit(MOTORDIR)
		io_write_analog(MOTOR,2800)
	}else if dirn < 0{
		io_set_bit(MOTORDIR)
		io_write_analog(MOTOR,2800)	
	}
}

func elev_get_floor_sensor_signal() int{
	if io_read_bit(SENSOR_FLOOR1) == 1{
		return 0
	}else if io_read_bit(SENSOR_FLOOR2) == 1{
		return 1
	}else if io_read_bit(SENSOR_FLOOR3) == 1{
		return 2
	}else if io_read_bit(SENSOR_FLOOR4) == 1{
		return 3
	}
	return -1	 
}

func elev_get_button_signal(button elev_button_type_t, floor int) bool{
	if io_read_bit(button_channel_matrix[floor][button]) != 0{
		return true	
	}else{
		return false
	}
} 

func elev_set_button_signal(button elev_button_type_t, floor int, value bool){
	if value == true{
		io_set_bit(lamp_channel_matrix[floor][button])
	}else{
		io_clear_bit(lamp_channel_matrix[floor][button])
	}
}

func elev_get_stop_signal() int{
	return io_read_bit(STOP)
}

func newInternalOrders(internalOrders [N_FLOORS]int) [N_FLOORS]int{
	for i:=0;i<N_FLOORS;i++{
		if elev_get_button_signal(BUTTON_COMMAND,i) == true{
			internalOrders[i] = 1	
		} 
	}
	return internalOrders
}
func clearInternalOrders(delete bool,floor int, internalOrders [N_FLOORS]int) [N_FLOORS]int{
	if floor != -1 && delete == true{
		internalOrders[floor] = 0
	}
	return internalOrders
}

func setInternalLights(internalOrders [N_FLOORS]int){
	for i:=0;i < N_FLOORS;i++{
		if internalOrders[i] == 1{
			elev_set_button_signal(BUTTON_COMMAND,i,true)	
		}else{
			elev_set_button_signal(BUTTON_COMMAND,i,false)
		}
	}
}

func newExternalOrders(externalOrders [N_FLOORS][2]int) [N_FLOORS][2]int{
	for i:=0;i<N_FLOORS-1;i++{
		if elev_get_button_signal(BUTTON_CALL_UP,i) == true{
			externalOrders[i][0] = 1
		}
	}
	for i:=1;i<N_FLOORS;i++{
		if elev_get_button_signal(BUTTON_CALL_DOWN,i) == true{
			externalOrders[i][1] = 1
		}
	}	
	return externalOrders
}

func clearExternalOrders(delete bool, floor int, externalOrders [N_FLOORS][2]int) [N_FLOORS][2]int{
	if floor != -1 && delete == true{
		externalOrders[floor][0] = 0
		externalOrders[floor][1] = 0
	}
	return externalOrders
}

func setExternalLights(externalOrders [N_FLOORS][2]int){
	for i:=0;i<N_FLOORS-1;i++{
		if externalOrders[i][0] == 1{
			elev_set_button_lamp(BUTTON_CALL_UP,i,true)
		}else{
			elev_set_button_lamp(BUTTON_CALL_UP,i,false)
		}
	}
	for i:=1;i<N_FLOORS;i++{
		if externalOrders[i][1] == 1{
			elev_set_button_lamp(BUTTON_CALL_DOWN,i,true)
		}else{
			elev_set_button_lamp(BUTTON_CALL_DOWN,i,false)
		}
	}
}

func InformationToNetworkUnit(internalOrders [N_FLOORS]int,externalOrders [N_FLOORS][2]int, ExternalOrdersToNetwork chan [N_FLOORS][2]int, InternalOrdersToNetwork chan [N_FLOORS]int){
	InternalOrdersToNetwork <- internalOrders
	ExternalOrdersToNetwork <- externalOrders
}
func States(FloorChan chan int){
	lastFloor := elev_get_floor_sensor_signal()
	for{
			if elev_get_floor_sensor_signal() != -1{
				lastFloor = elev_get_floor_sensor_signal()
			}					
			FloorChan <- lastFloor

	}	
}

func stopAtFloor(order []int,lastFloor int)bool{
	stop := false
	for i:= 0; i < len(order);i++{
		if order[i] == 1{
			if i == lastFloor{
				stop = true
			}	
		}
	} 
	return stop
}

func doorTimer(DirectionChan chan int, ExecuteListChan chan []int){	
	doorTimer:=time.Now().Add(time.Second*2).UnixNano()/int64(time.Millisecond)
	for{ 
		<- DirectionChan
		<- ExecuteListChan
		if time.Now().UnixNano()/int64(time.Millisecond) > doorTimer{
		break
		}
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func read(filename string) ([4]int, [4][2]int) {

    if _, err := os.Stat(filename); os.IsNotExist(err) {
        Println("File does not exist at location. Making new queue file.")
        os.Create(filename)
    }

    type order_struct struct{
    	internal [4]int
    	external [4][2]int
    }

    dat, err := ioutil.ReadFile(filename)
    check(err)

    Orders := new(order_struct)

    json.Unmarshal(dat, &Orders)

    var internalOrders [4]int
    var externalOrders [4][2]int

    internalOrders = Orders.internal
    externalOrders = Orders.external

    return internalOrders, externalOrders
}

func write(filename string, internalOrders [4]int, externalOrders [4][2]int) {

    if _, err := os.Stat(filename); os.IsNotExist(err) {
    	Println("File does not exist at location. Making new queue file.")
        os.Create(filename)
    }

    //Println("Hei")

    type order_struct struct{
    	internal [4]int
    	external [4][2]int
    }
    
    Orders := order_struct{internal:internalOrders, external: externalOrders}
    Println(Orders)
    //Println(Orders)

    b,_ := json.Marshal(Orders)

    //Println(b)

    err := ioutil.WriteFile(filename, b, 0644)
    check(err)
}

func ExecuteOrders(ExecuteListChan chan []int, DirectionChan chan int,LastStopChannel chan int){
	go States(FloorChan)
	floor := elev_get_floor_sensor_signal()
	for{
		direction :=<- DirectionChan
		order :=<- ExecuteListChan				
		switch{
		case elev_get_floor_sensor_signal() != -1: //elevator is at a floor
			floor = elev_get_floor_sensor_signal()
			if stopAtFloor(order,floor) == true{
				DeleteOrderChan <- true
				elev_set_motor_direction(DIRN_STOP)
				
				
				elev_set_door_open_lamp(true)
				doorTimer(DirectionChan,ExecuteListChan)
				elev_set_door_open_lamp(false)
			}
			if direction > 0 && elev_get_floor_sensor_signal() != 3{
				elev_set_motor_direction(DIRN_UP)
			}
			if direction < 0 && elev_get_floor_sensor_signal() != 0{
				elev_set_motor_direction(DIRN_DOWN)
			}
			if direction == 0{
				elev_set_motor_direction(DIRN_STOP)
				LastStopChannel <- elev_get_floor_sensor_signal()
			}

		case elev_get_floor_sensor_signal() == -1:			
					
		}
	}
}


func lightsAndOrders(internalOrders [N_FLOORS]int,DirnChan chan int,GlobalExternalOrdersChannel chan [4][2]int,ExecutedChannel chan bool){	
	GlobalExternalOrders := [N_FLOORS][2]int{{0,0},{0,0},{0,0},{0,0}}
	//externalOrders := [N_FLOORS][2]int{{0,0},{0,0},{0,0},{0,0}}

			externalOrders := [N_FLOORS][2]int{{0,0},{0,0},{0,0},{0,0}}
			externalOrders = newExternalOrders(externalOrders) 
			InformationToNetworkUnit(internalOrders,externalOrders,ExternalOrdersToNetwork,InternalOrdersToNetwork)
			setExternalLights(GlobalExternalOrders)
			internalOrders = newInternalOrders(internalOrders)
			setInternalLights(internalOrders)	
			elev_set_floor_indicator(elev_get_floor_sensor_signal())
	for{

		select{

		case delete :=<-DeleteOrderChan:
			floor := elev_get_floor_sensor_signal()
			internalOrders = clearInternalOrders(delete,floor,internalOrders)
			ExecutedChannel <- delete
			externalOrders = clearExternalOrders(delete,floor,externalOrders)

		case temp :=<- GlobalExternalOrdersChannel:

			GlobalExternalOrders = temp

		default:

			externalOrders = newExternalOrders(externalOrders) 
			InformationToNetworkUnit(internalOrders,externalOrders,ExternalOrdersToNetwork,InternalOrdersToNetwork)
			internalOrders = newInternalOrders(internalOrders)
			setInternalLights(internalOrders)	
			elev_set_floor_indicator(elev_get_floor_sensor_signal())
			setExternalLights(GlobalExternalOrders)	
		}
	}
}

func runElevator(){
	elev_init()
	elev_set_motor_direction(DIRN_STOP)
	
	go ExecuteOrders(ExecuteListChan,DirectionChan,LastStopChannel)
}

func Elevator(){

	type order_struct struct{
    	Internal [4]int
    	External [4][2]int
    }

	internalOrders := [N_FLOORS]int{0,0,0,0}
	externalOrders := [N_FLOORS][2]int{{0,0},{0,0},{0,0},{0,0}}

	filename := "order_lists"
	_, err := os.Stat(filename)

	if!(os.IsNotExist(err)){

		dat, err := ioutil.ReadFile(filename)
    	check(err)

    	orders := new(order_struct)

    	json.Unmarshal(dat, &orders)

    	internalOrders = orders.Internal
    	externalOrders = orders.External

    	Println(internalOrders)
    	Println(externalOrders)

	}else{
		write(filename, internalOrders, externalOrders)
	}

	//Println(read(filename))

	go runElevator()
	go lightsAndOrders(internalOrders, DirnChan,GlobalExternalOrdersChannel,ExecutedChannel)

	Println("Elevator")
}



