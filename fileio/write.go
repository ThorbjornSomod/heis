package main

import (
    "bufio"
    "io"
    "os"
)

func write(filename string){

// open output file
    fo, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    // make a write buffer
    w := bufio.NewWriter(fo)

    for{
    	// write a chunk
        if _, err := w.Write(buf[:n]); err != nil {
            panic(err)
        }
    }

    if err = w.Flush(); err != nil {
        panic(err)
    }

}

func main(){
	write("output.txt")
}