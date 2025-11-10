package main

type Job interface {
	ID int
	Execute func(int)
}

type Worker interface {
	ID int
	JobChannel chan Job
	WorkerPool chan chan Job
	Quit chan bool
}

func NewWorker(id int, pool chan chan Job) Worker {
	return Worker{
		ID: id,
		JobChannel: make(chan Job),
		WorkerPool: pool,
		Quit: make(chan bool),
	}
}

func (w *Worker) Run(ctx context.Context) {
	for {
		w.WorkerPool <- w.JobChannel
		select {
		case job := <-w.JobChannel:
			job.Execute(w.ID)
		case w.Quit <- true:
			return
		case <-ctx.Done():
			return
		}
	}
}