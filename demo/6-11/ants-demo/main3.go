package main

import (
	"fmt"
	"sync"
)

type NamedWorkerPool struct {
	jobs       chan Job
	waitGroup  sync.WaitGroup
	numWorkers int
}

type Job struct {
	Name string
	Task func()
}

func NewNamedWorkerPool(numWorkers int) *NamedWorkerPool {
	return &NamedWorkerPool{
		jobs:       make(chan Job),
		numWorkers: numWorkers,
	}
}

func (p *NamedWorkerPool) Start() {
	for i := 0; i < p.numWorkers; i++ {
		go func() {
			for job := range p.jobs {
				fmt.Printf("Worker %s is starting\n", job.Name)
				job.Task()
				fmt.Printf("Worker %s is done\n", job.Name)
				p.waitGroup.Done()
			}
		}()
	}
}

func (p *NamedWorkerPool) Submit(name string, task func()) {
	p.waitGroup.Add(1)
	p.jobs <- Job{Name: name, Task: task}
}

func (p *NamedWorkerPool) Wait() {
	p.waitGroup.Wait()
}

func main() {
	pool := NewNamedWorkerPool(3)

	pool.Start()

	pool.Submit("Job 1", func() {
		fmt.Println("Job 1 is running")
	})

	pool.Submit("Job 2", func() {
		fmt.Println("Job 2 is running")
	})

	pool.Submit("Job 3", func() {
		fmt.Println("Job 3 is running")
	})

	pool.Wait()
}
