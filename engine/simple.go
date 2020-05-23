package engine

import (
	"fmt"
	"log"
	"time"
)

func SimpleRun(timeDelay int, seeds ...Request) {
	// 定义一个任务队列
	var requests []Request
	// 将种子任务放入队列中
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) != 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Parsing URL:%s", r.CurURL)
		parserResult, err := Worker(r)
		if err != nil {
			log.Print(err)
			continue
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got item %v, res: %d", item, len(requests))
		}
		fmt.Println(len(requests))
		// 设置每次请求的延时时间，防止请求过于频繁被拉黑
		time.Sleep(time.Second * time.Duration(timeDelay))
	}
}
