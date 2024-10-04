package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("sync with test")
}

// using sync package for concurent task
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
	return &Counter{}
}
