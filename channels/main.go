package main

import (
	"fmt"
	"sync"
	"time"
)

// Notes:
// - use channels to transfer data between goroutines

type worker struct {
	numbersChannel chan int
}

func (w *worker) DoWork(n int) {
	time.Sleep(time.Second) // simulates work
	w.numbersChannel <- n   // sends number to the channel
}

func (w *worker) Print() {
	for n := range w.numbersChannel {
		fmt.Printf("n = %d\n", n)
	}
}

func (w *worker) Close() {
	close(w.numbersChannel)
}

var wg sync.WaitGroup

func main() {
	w := worker{
		numbersChannel: make(chan int),
	}

	for i := 0; i < 100; i++ {
		wg.Add(1) // adds counter +1

		// spawns a goroutine
		go func(i int) {
			defer wg.Done() // decrements counter -1, and is deferred until the function ends
			w.DoWork(i)     // sends number to the chanenl
		}(i)
	}

	// waits for all goroutines to finish
	go func() {
		wg.Wait() // waits for the counter to reach 0
		w.Close()
	}()

	w.Print()
	fmt.Println("Done")
}

// more notes:
// by default, channels have no space inside and so you cannot temporarily store a value in it, it is more like a portal
// (when unbuffered)
// you need to simultaneously have 1 thread to send data to channel
// and another to receive:
// -----------------------
//	numbersChannel := make(chan int)
//
//	go func() {
//		numbersChannel <- 555
//	}()
//
//	n := <-numbersChannel
//
//	fmt.Printf("n = %d\n", n)
// -----------------------
// (when buffered)
// you can send/receive under 1 thread:
// numbersChannel := make(chan int, 1)
//
//	numbersChannel <- 555
//
//	n := <-numbersChannel
//
//	fmt.Printf("n = %d\n", n)
// -----------------------
