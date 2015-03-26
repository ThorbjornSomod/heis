package elevator

import(
		//."fmt"
		//"time"
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
func clearInternalOrders(dirn int, floor int, internalOrders [N_FLOORS]int) [N_FLOORS]int{
	if dirn == 0 && floor != -1{
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

func newExternalOrders(externalOrders [N_FLOORS][N_BUTTONS]int) [N_FLOORS][N_BUTTONS]int{
	for i:=0;i<N_FLOORS-1;i++{
		if elev_get_button_signal(BUTTON_CALL_UP,i) == true{
			externalOrders[i][0] = BUTTON_CALL_UP
			externalOrders[i][2] = 1
		}
	}
	for i:=1;i<N_FLOORS;i++{
		if elev_get_button_signal(BUTTON_CALL_DOWN,i) == true{
			externalOrders[i][1] = BUTTON_CALL_DOWN
			externalOrders[i][2] = 1
		}
	}	
	return externalOrders
}

func clearExternalOrders(dirn int, floor int, externalOrders [N_FLOORS][N_BUTTONS]int) [N_FLOORS][N_BUTTONS]int{
	if dirn == 0 && floor != -1{
		externalOrders[floor][0] = -1
		externalOrders[floor][1] = -1
		externalOrders[floor][2] = -1
	}
	return externalOrders
}

func setExternalLights(externalOrders [N_FLOORS][N_BUTTONS]int){
	for i:=0;i<N_FLOORS-1;i++{
		if externalOrders[i][0] == 0{
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

func OrdersToNetworkUnit(internalOrders [N_FLOORS]int,externalOrders [N_FLOORS][N_BUTTONS]int, ExternalOrdersToNetwork chan [N_FLOORS][N_BUTTONS]int, InternalOrdersToNetwork chan [N_FLOORS]int){
	InternalOrdersToNetwork <- internalOrders
	ExternalOrdersToNetwork <- externalOrders
}

/*
func OrdersFromNetworkUnit(ExecuteFromNetwork chan []int, NextFloor chan []int){ //Mottar execute list og sender første element på lista til å bli executed
	execute := <-ExecuteFromNetwork
	NextFloor <- execute[0]

}
*/

func lightsAndOrders(internalOrders [N_FLOORS]int, externalOrders [N_FLOORS][N_BUTTONS]int,dirn int){
	

	for{
		internalOrders = newInternalOrders(internalOrders)
		externalOrders = newExternalOrders(externalOrders)
		setInternalLights(internalOrders)
		setExternalLights(externalOrders)
		internalOrders = clearInternalOrders(dirn,elev_get_floor_sensor_signal(),internalOrders)
		externalOrders = clearExternalOrders(dirn,elev_get_floor_sensor_signal(),externalOrders)
		elev_set_floor_indicator(elev_get_floor_sensor_signal())
		OrdersToNetworkUnit(internalOrders,externalOrders,ExternalOrdersToNetwork,InternalOrdersToNetwork)
	}
	
}
func runElevator(){
	elev_init()
	elev_set_motor_direction(DIRN_STOP)

	for{
		//Safety feature; turn elevator when it reaches boundary
		if elev_get_floor_sensor_signal() == N_FLOORS -1{
			elev_set_motor_direction(DIRN_STOP)
		}else if elev_get_floor_sensor_signal() == 0{
			elev_set_motor_direction(DIRN_STOP)
		}
		if elev_get_stop_signal() != 0{
			elev_set_motor_direction(DIRN_STOP)
			break
		}
	}		
}

func Elevator(){
	dirn := DIRN_STOP
	internalOrders := [N_FLOORS]int{0,0,0,0}
	externalOrders := [N_FLOORS][N_BUTTONS]int{{-1,-1,-1},{-1,-1,-1},{-1,-1,-1},{-1,-1,-1}}
	
	go runElevator()
	go lightsAndOrders(internalOrders,externalOrders,dirn)

	

}

