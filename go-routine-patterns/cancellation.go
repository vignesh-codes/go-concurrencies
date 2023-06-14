package main

import (
	"context"
	"fmt"
	"time"
)

/*
Here the work keeps doing and will be cancelled after 3 seconds until then
the work continue to happen inside for loop
*/
func initCancellation() {
	fmt.Println("Initializing cancellation pattern")

	ctx, cancel := context.WithCancel(context.Background())
	// provides with cancellation support for this ctx
	go doWork1(ctx)
	time.Sleep(time.Second * 3)

	cancel()

	time.Sleep(time.Second * 1)
}

func doWork1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // ctx.done receives a value if cancelled
			fmt.Println("work is canceled")
			return
		default:
			fmt.Println("working is continueing")
		}

		time.Sleep(time.Second * 1)
	}
}
