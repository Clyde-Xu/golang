package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	dispatcher := NewDispatcher(3)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dispatcher.Run(ctx)

	for i := 1; i <= 10; i++ {
		job := Job{
			ID: i,
			Execute: func(id int) {
				fmt.Printf("Worker processed job #%d\n", id)
				time.Sleep(time.Millisecond * 500)
			},
		}
		dispatcher.Submit(job)
	}

	// time.Sleep(3 * time.Second)
	fmt.Println("Shutting down...")
	dispatcher.Stop()

}
