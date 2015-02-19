package main 

import (
		//"net"
		//"bufio"
		."fmt"
		"runtime"
		"time"
		."./network"
		"io/ioutil"
		"os"
		//"os/exec"
		."encoding/json"
)
 

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	BroadcastIP, WelcomePort,client := Init()
	errorChan := make(chan bool)
	IPchannel := make(chan string)
	go AddNewClient(BroadcastIP,WelcomePort,client,IPchannel,errorChan)
	Println(client)

	
	
		
	if client == "slave"{
		Println("saldhals")
		for{	
			Println("dlshjlsdfhgdfshg")
			a := <-errorChan
			Println(a)
			if a{
				client = "master"
				Println("Hei")				
				return

			}else{
				Println("aa")
			}
			time.Sleep(1000*time.Millisecond)	
		}
	}
				
	if client == "master"{
		//c := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run main.go")
		//c.Run()

		if _, err := os.Stat("backup.txt"); os.IsNotExist(err){
			os.Create("backup.txt")
			temp := 0				
			n, _ := Marshal(temp)
			ioutil.WriteFile("backup.txt", n, 0644)
		}
			
		for {
			f,_ := ioutil.ReadFile("backup.txt")
			var b int
			Unmarshal(f,&b)
			counter := b
			Println(counter)
			counter++
			n, _ := Marshal(counter)
			ioutil.WriteFile("backup.txt", n, 0644)
			time.Sleep(1000*time.Millisecond)
		}
	
	}
}