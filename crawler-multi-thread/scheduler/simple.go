package scheduler

import "golang-demos/crawler-multi-thread/engine"

// SimpleScheduler struct
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// WorkerChan create a worker chan
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// WorkerReady inform a worker is ready
func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
}

// Run assign request to worker
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// Submit submit a request to worker channel
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}
