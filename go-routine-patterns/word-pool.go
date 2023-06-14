package main

import (
	"fmt"
	"sync"
)

/*
go routines are spawned from a job queue.
The workers will listen to this shared job queue and will process them
*/
func initWorkerPool() {
	fmt.Println("initializing worker-pool pattern")
	numWorkers := 3
	jobQueue := make(chan int)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for job := range jobQueue {
				process(job)
			}
		}()
	}

	go func() {
		for i := 1; i <= 10; i++ {
			jobQueue <- i
		}
		close(jobQueue)
	}()

	go func() {
		wg.Wait()
		done <- true
	}()

	<-done
}

func process(job int) {
	fmt.Println("processing job ", job)
}
