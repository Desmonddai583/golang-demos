package scheduler

import "golang-demos/crawler-multi-thread/engine"

// QueuedScheduler struct
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

// WorkerChan create a worker chan
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// Submit submit a request to worker channel
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// WorkerReady inform a worker is ready
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

// Run assign request to worker
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
