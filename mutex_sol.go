package main 

import (
	."fmt"
	"runtime"
	// "time"
)

var i int




func function1(chn chan int, ch_inc chan string){
	for x := 0; x < 10000; x++ {
		<- chn
		i++
		chn <-1
		Println("func1", i)
	}
}

func function2(chn chan int, ch_dec chan string){
	for x := 0; x < 10000; x++ {
		<- chn
		i--
		chn <-1
		Println("func2", i)
	}
}

func main(){
	chn := make(chan int)
	ch_inc := make(chan string)
	ch_dec := make(chan string)
	i = 0
	
	runtime.GOMAXPROCS(runtime.NumCPU())

	go function1(chn, ch_inc)
	go function2(chn, ch_dec)

	chn <- 1

	// time.Sleep(1000*time.Millisecond)

	// Println(i)
	Println(<- ch_inc)
	Println(<- ch_dec)
}