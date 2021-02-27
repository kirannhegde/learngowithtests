package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//Channel to store the user input
	userinput := make(chan []byte)
	//Channel on which the current time gets sent after the 30ms is elapsed
	timeout := time.After(30 * time.Second)
	go gatherUserInput(userinput)
	select {
	case <-timeout:
		fmt.Printf("Timed out")
		os.Exit(0)
	case data := <-userinput:
		os.Stdout.Write(data)
	}
}

func gatherUserInput(in chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			in <- data
		}

	}
}
