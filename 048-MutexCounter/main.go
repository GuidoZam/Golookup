package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mutex sync.Mutex
	value  map[string]int
}

// Increment increments the counter for the given key.
func (c *SafeCounter) Increment(key string) {
	defer c.mutex.Unlock()

	c.mutex.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.value[key]++
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mutex.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mutex.Unlock()
	return c.value[key]
}

func main() {
	c := SafeCounter{value: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Increment("somekey")
	}

	// The sleep is to count up to 1000
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

