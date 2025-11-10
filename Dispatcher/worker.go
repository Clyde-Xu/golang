package main

type Dispatcher struct {
	WorkerPool chan chan Job
}