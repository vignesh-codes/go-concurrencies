package main

import (
	"fmt"
	"time"
)

/*
This pattern helps you to pass values across different stages and makes sure a
specific operation is performed by the stage before it
*/
func initPipeline() {
	fmt.Println("Initializing pipeline pattern")
	input := make(chan int)
	output1 := make(chan int)
	output2 := make(chan int)

	// Start the pipeline stages
	go stage1(input, output1)
	// output2 is written by stage1 which is taken as input by stage2
	go stage2(output1, output2)
	// output2 is written by stage2 and is taken as input by stage3
	go stage3(output2)
	for i := 0; i <= 3; i++ {
		input <- i
		time.Sleep(time.Second * 1)
	}
	// we always close after writing to channel
	close(input)
	for range output2 {
	}
	fmt.Println("Pipeline completed")
}

func stage1(input <-chan int, output chan<- int) {
	for value := range input {
		result := value * 2
		fmt.Printf("stage1: %d -> %d\n", value, result)
		output <- result
	}
	close(output)
}

func stage2(input <-chan int, output chan<- int) {
	for value := range input {
		result := value * 3
		fmt.Printf("stage2: %d -> %d\n", value, result)
		output <- result
	}
	close(output)
}

func stage3(input <-chan int) {
	for value := range input {
		result := value * 4
		fmt.Printf("stage3: %d -> %d\n", value, result)
	}
}
