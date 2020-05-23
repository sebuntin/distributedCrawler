package main

import (
	. "doubanparser"
	"engine"
	"persist"
	"scheduler"
	"worker"
)

const SEEDURL = "https://book.douban.com/tag/?view=type&icn=index-sorttags-all"
const HeadURL = "https://book.douban.com"

//const Prefix = "https://book.douban.com"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		NumWorkers:       10,
		ItemChan:         persist.ItemSaver(),
		RequestProcessor: worker.CreatProcessor(DoubanParseFuncConversion{}),
	}
	e.Run(engine.Request{
		CurURL: SEEDURL,
		Args:   []string{SEEDURL},
		Parser: engine.NewFuncParser("ParseBookTag",
			[]string{SEEDURL}, ParseBookTag),
	})
}
