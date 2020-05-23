package scheduler

import "engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workChan
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {}

func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workChan <- r }()
}
