package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Notes:
// - use channels to transfer data between goroutines

var wg sync.WaitGroup

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	dataChan := make(chan int)

	go func() {
		for i := 0; i <= 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				randomNum := DoWork()
				dataChan <- randomNum
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n = %d\n", n)

	}
}

// by default, channels have no space inside and so you cannot temporarily store a value in it, it is more like a portal
// (when unbuffered)
// you need to simultaneously have 1 thread to send data to channel
// and another to receive:
// -----------------------
//	dataChan := make(chan int)
//
//	go func() {
//		dataChan <- 555
//	}()
//
//	n := <-dataChan
//
//	fmt.Printf("n = %d\n", n)
// -----------------------
// (when buffered)
// you can send/receive under 1 thread:
// dataChan := make(chan int, 1)
//
//	dataChan <- 555
//
//	n := <-dataChan
//
//	fmt.Printf("n = %d\n", n)
// -----------------------
