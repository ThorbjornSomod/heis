package channels
import (
//"net"
//."fmt"
//"runtime"
//"time"
)

var MasterIsAlive = make(chan string)
var AliveMessage = make(chan string)
var IPlistChan = make(chan []string)
var ExecuteListChan = make(chan []int)
var IPchan = make(chan string)
var NextFloor = make(chan []int)
var InternalOrdersToNetwork = make(chan [4]int)
var ExternalOrdersToNetwork = make(chan [4][3]int)
var DirnChan = make(chan int)
