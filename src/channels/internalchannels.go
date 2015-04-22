package channels
import (
//"net"
//."fmt"
//"runtime"
//"time"
)


var ExecuteListChan = make(chan []int)
var IPchan = make(chan string)
var NextFloor = make(chan []int)
var InternalOrdersToNetwork = make(chan [4]int)
var ExternalOrdersToNetwork = make(chan [4][2]int)
var DirectionChan = make(chan int)
var FloorChan = make(chan int)
var DirnChan = make(chan int)
var DirChan = make(chan int)