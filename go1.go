package main 

import (
	."fmt"
	"runtime"
	"time"
)

var i int

func function1(){
	for x := 0; x < 1000000; x++ {
		i++
	}
}

func function2(){
	for x := 0; x < 1000000; x++ {
		i--
	}
}

func main(){

	i = 0

	runtime.GOMAXPROCS(runtime.NumCPU())

	go function1()
	go function2()

	time.Sleep(100*time.Millisecond)

	Println(i)

}