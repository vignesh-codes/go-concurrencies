package main

import (
	"fmt"
	"time"
)

/*
This pattern allows you to spawn a go routine with a timeout.
If timeout passes, the function exits
*/
func initTimeout() {
	fmt.Println("Initializing timeout pattern")

	result := completeWithinTime(doWork, time.Second*3)
	fmt.Println("job completed? ", result)
}

func completeWithinTime(work func() bool, timeout time.Duration) bool {
	done := make(chan bool)
	go func() {
		done <- work()
	}()

	select {
	case out := <-done:
		return out
	case <-time.After(timeout):
		return false
	}

}

func doWork() bool {
	time.Sleep(time.Second * 2)
	return true
}
