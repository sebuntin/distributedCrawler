package scheduler

import "engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkReady(c chan engine.Request) {
	q.workerChan <- c
}

func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		requestQ := make([]engine.Request, 0, 1000)
		workerQ := make([]chan engine.Request, 0, 1000)
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 当request和worker同时存在时
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
