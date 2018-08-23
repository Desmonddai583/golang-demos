package scheduler

import "golang-demos/crawler-multi-thread/engine"

// SimpleScheduler struct
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// ConfigureMasterWorkerChan set channel to master worker
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

// Submit submit a request to worker channel
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}
