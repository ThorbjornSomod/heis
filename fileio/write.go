package main

import (
    "io/ioutil"
    "encoding/json"
    "os"
    //."fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func write(filename string, internalOrders []int, externalOrders []int) {

    if _, err := os.Stat(filename); os.IsNotExist(err) {
        os.Create(filename)
    }

    var orders [2][]int

    orders[0] = internalOrders
    orders[1] = externalOrders

    b,_ := json.Marshal(orders)

    err := ioutil.WriteFile(filename, b, 0644)
    check(err)
}

func main(){
    internalOrders := []int{1,2,3,4,7,10}
    externalOrders := []int{3,2,1,9,7,11}
	write("input.txt",internalOrders, externalOrders)
}