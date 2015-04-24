package main

import (
    "io/ioutil"
    ."fmt"
    "encoding/json"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func read(filename string) ([]int, []int) {

    if _, err := os.Stat(filename); os.IsNotExist(err) {
        Println("File does not exist at location. Making new queue file.")
        os.Create(filename)
    }

    dat, err := ioutil.ReadFile(filename)
    check(err)

    var b [2][]int
    json.Unmarshal(dat, &b)

    internalOrders := b[0]
    externalOrders := b[1]

    return internalOrders, externalOrders
}

func main(){
    Println(read("input.txt"))
}