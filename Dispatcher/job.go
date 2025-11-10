package main

import (
	"context"
	"fmt"
)

type Job struct {
	ID      int
	Execute func(int)
	Done    chan bool
}

type Worker struct {
	ID         int
	JobChannel chan Job
	WorkerPool chan chan Job
	Quit       chan bool
}

func NewWorker(id int, pool chan chan Job) Worker {
	return Worker{
		ID:         id,
		JobChannel: make(chan Job),
		WorkerPool: pool,
		Quit:       make(chan bool),
	}
}

func (w *Worker) Start(ctx context.Context) {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				fmt.Printf("Worker %d received job #%d\n", w.ID, job.ID)
				job.Execute(job.ID)
			case w.Quit <- true:
				return
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.Quit)
}
