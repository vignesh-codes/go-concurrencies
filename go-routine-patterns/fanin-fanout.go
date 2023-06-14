package main

import (
	"fmt"
	"sync"
	"time"
)

/*
This pattern comes in handy when you have a independent producer which has to be processed
before consuming.
Here multiple go routine can produce which will be processed the same way and consumed
by only one go routine.
*/
func initFanInFanOut() {
	fmt.Println("Initializing fan-in fan-out pattern")

	input := make(chan int)
	output := make(chan int)
	var wg sync.WaitGroup

	go produceData(input, 10)
	numWorkers := 1
	for i := 0; i <= numWorkers; i++ {
		wg.Add(1)
		go processInput(input, output, &wg)
	}
	go consumeData(output)

	wg.Wait()
}

// gets the input sent by produce data and process it to output
func processInput(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range input {
		output <- value * 10
	}
}

// produces the input data to input chan which is listened by processInput
func produceData(input chan<- int, count int) {
	defer close(input)
	for i := 0; i <= count; i++ {
		time.Sleep(1 * time.Second)

		input <- i
	}
}

// consumes the output that is received
func consumeData(output <-chan int) {
	for value := range output {
		fmt.Println("consumed val ", value)
	}
}
