package engine

import (
	"log"
)

type Scheduler interface {
	Submit(Request)
	WorkChan() chan Request
	WorkReady(chan Request)
	Run()
}

type Processor func(Request) (ParserResult, error)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	ItemChan         chan interface{}
	NumWorkers       int
	RequestProcessor Processor
}

func (e ConcurrentEngine) Run(seed ...Request) {
	//in := make(chan Request, 10)
	out := make(chan ParserResult)
	//e.Scheduler.ConfigureMasterWorkChan(in)
	e.Scheduler.Run()
	// 将种子请求放入RequestChan中
	for _, r := range seed {
		if isDuplicate(r.CurURL) {
			e.Scheduler.Submit(r)
		}
	}
	// 创建worker
	for i := 0; i < e.NumWorkers; i++ {
		e.creatWorker(e.Scheduler, out)
	}
	//var count int
	for {
		result := <-out
		for _, item := range result.Items {
			//count ++
			//log.Printf("item: %v, count: %d\n", item, count)
			// 这里需要注意不能写成
			// go func() {c <- item} ()
			// 由于闭包函数内使用的外部变量,会受到外部影响。
			go func(c chan interface{},
				item interface{}) {
				c <- item
			}(e.ItemChan, item)
		}

		for _, r := range result.Requests {
			if isDuplicate(r.CurURL) {
				e.Scheduler.Submit(r)
			}
		}
	}
}

func (e *ConcurrentEngine) creatWorker(scheduler Scheduler, out chan ParserResult) {
	in := scheduler.WorkChan()
	go func() {
		for {
			scheduler.WorkReady(in)
			r := <-in
			// here to call crawler rpc service
			result, err := e.RequestProcessor(r)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			//time.Sleep(time.Second * 1)
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return false
	}
	visitedUrls[url] = true
	return true
}
