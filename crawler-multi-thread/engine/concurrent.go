package engine

import (
	"golang-demos/crawler-multi-thread/model"
	"log"
)

// ConcurrentEngine struct
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// Scheduler interface
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// ReadyNotifier interface
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run fetch and parse flow
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.URL) {
			log.Printf("Duplicate request: "+
				"%s", r.URL)
			continue
		}
		e.Scheduler.Submit(r)
	}

	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got profile #%d: %v", profileCount, item)
				profileCount++
			}
		}

		for _, request := range result.Requests {
			if isDuplicate(request.URL) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
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

var visitedURLs = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedURLs[url] {
		return true
	}

	visitedURLs[url] = true
	return false
}
