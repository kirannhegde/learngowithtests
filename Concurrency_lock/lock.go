package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		go processLock(lock, i)
	}
	select {
	case status := <-lock:
		if status {
			fmt.Printf("Lock deactivated by go routine")
		} else {
			fmt.Printf("All processing is complete")
			return
		}
	default:
		fmt.Printf("Go routines are still processing.")

	}

}

func processLock(in chan bool, number int) {
	in <- true
	fmt.Printf("Lock activated by go routine:%d", number)
	time.Sleep(5 * time.Second)
	fmt.Printf("Processing has stopped")
}
