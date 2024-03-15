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
	mutex     sync.Mutex
}

func (c *counter) add(key string, num int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.numberMap[key] = num
}

var wg sync.WaitGroup

func main() {
	c := counter{
		numberMap: make(map[string]int),
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			c.add(key, i)
		}(i)
	}
	wg.Wait()

	for key, num := range c.numberMap {
		fmt.Printf("Key: %s, Value: %d\n", key, num)
	}
}
