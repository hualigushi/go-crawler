package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		creatWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)

	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func creatWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
