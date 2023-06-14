package main

import (
	"fmt"
	"sync"
)

/*
Semaphore can be used to limit the number of go routines at a time.
Here we define concurrency as 3.
we initiate our go routine and assign empty struct to semaphore
once finished, we release the empty struct from semaphore
*/
func initSemaphore() {
	fmt.Println("initializing semaphore pattern")
	concurrency := 3
	semaphore := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(jobId int) {
			semaphore <- struct{}{}
			defer func() {
				<-semaphore
				wg.Done()
			}()
			processJob(jobId)
		}(i)
	}

	wg.Wait()
}

func processJob(job int) {
	fmt.Println("processing job ", job)
}
