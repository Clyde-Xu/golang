package main

import (
	"context"
	"fmt"
)

type Dispatcher struct {
	WorkerPool chan chan Job
	JobQueue   chan Job
	Workers    []Worker
	Quit       chan bool
}

func NewDispatcher(workers int) *Dispatcher {
	pool := make(chan chan Job, workers)
	for i := 0; i < workers; i++ {
		pool <- make(chan Job)
	}
	return &Dispatcher{
		WorkerPool: pool,
		JobQueue:   make(chan Job),
		Workers:    make([]Worker, workers),
		Quit:       make(chan bool),
	}
}

func (d *Dispatcher) Submit(job Job) {
	d.JobQueue <- job
}

func (d *Dispatcher) Run(ctx context.Context) {
	for i := range d.Workers {
		worker := NewWorker(i, d.WorkerPool)
		d.Workers[i] = worker
		worker.Start(ctx)
	}
	go func() {
		fmt.Println("Starting dispatcher...")
		for {
			select {
			case job := <-d.JobQueue:
				go func(job Job) {
					jobChan := <-d.WorkerPool
					jobChan <- job
				}(job)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (d *Dispatcher) Stop() {
	for _, w := range d.Workers {
		w.Stop()
	}
	close(d.JobQueue)
	close(d.WorkerPool)
}
