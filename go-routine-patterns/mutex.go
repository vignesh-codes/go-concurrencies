package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.value++
}

func (c *Counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.value
}

/*
Prevents concurrent read and write errors
*/
func initMutex() {
	fmt.Println("Initializing Mutex Locks Pattern")
	counter := Counter{}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	fmt.Println("Final counter value:", counter.GetValue())
}
