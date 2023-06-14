package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	numChan := make(chan int)
	stringChan := make(chan string)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			numChan <- i
		}

		close(numChan)
	}()

	go func() {
		defer wg.Done()
		for i := 'a'; i <= 'z'; i++ {
			// time.Sleep(5 * time.Second)
			stringChan <- string(i)
		}

		close(stringChan)
	}()
	res := ""
	for num := range numChan {
		str, ok := <-stringChan
		if !ok {
			break
		}

		curatedStr := fmt.Sprintf("%d%s", num, str)
		res += curatedStr
		// fmt.Println(curatedStr)
	}
	fmt.Println(res)
}
