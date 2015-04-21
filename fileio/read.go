package main

import (
    "bufio"
    "io"
    "os"
    ."fmt"
)

func read(filename string) {
    // open input file
    fi, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    // make a read buffer
    r := bufio.NewReader(fi)

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        
        if n == 0 {
            break
        }   
    }
    Println(string(buf[10]))
}

func main(){
    read("input.txt")
}