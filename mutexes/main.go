package main

import (
	"fmt"
	"sync"
)

// Notes: (Mutual exclusions)
// - reach for this when you do not need to worry about communication between goroutines
// - and when you want to safely write to a single spot in memory (ensuring this space in mem is only written to by 1 thing)
// - by locking, doing the write, then unlocking

type counter struct {
	numberMap map[string]int
	mu        sync.Mutex
}

func (c *counter) add(num int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.numberMap["key"] = num
}

func main() {
	c := counter{numberMap: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.add(i)
		}(i)
	}

	wg.Wait()
	fmt.Printf("%d\n", c.numberMap["key"])
}
