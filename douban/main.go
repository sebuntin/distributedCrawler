package main

import (
	. "crawlerService"
	. "doubanparser"
	"engine"
	"flag"
	"fmt"
	"log"
	"persist"
	"scheduler"
	"strings"
	"worker"

	"google.golang.org/grpc"
)

const SEEDURL = "https://book.douban.com/tag/?view=type&icn=index-sorttags-all"
const HeadURL = "https://book.douban.com"

//const Prefix = "https://book.douban.com"

var (
	itemSaverHost = flag.Int("item_saver", 0,
		"Item Saver Host")
	workerHosts = flag.String("worker", "",
		"worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	fmt.Println(*itemSaverHost)
	servicePool := createClientPool(
		strings.Split(*workerHosts, ","))
	e := engine.ConcurrentEngine{
		Scheduler:  &scheduler.QueuedScheduler{},
		NumWorkers: 100,
		ItemChan:   persist.ItemSaver(*itemSaverHost),
		RequestProcessor: worker.CreatProcessor(servicePool,
			DoubanParseFuncConversion{}),
	}
	e.Run(engine.Request{
		CurURL: SEEDURL,
		Args:   []string{SEEDURL},
		Parser: engine.NewFuncParser("ParseBookTag",
			[]string{SEEDURL}, ParseBookTag),
	})
}

// 创建rpc client池
func createClientPool(hosts []string) chan CrawlerServiceClient {
	clients := make([]CrawlerServiceClient, 0)
	// 建立rpc client
	for i := range hosts {
		PORT := fmt.Sprintf(":%s", hosts[i])
		// grpc dial
		conn, err := grpc.Dial(PORT, grpc.WithInsecure())
		if err != nil {
			log.Printf("Create client pool: "+
				"error connecting to %s, %v\n", PORT, err)
			continue
		}
		client := NewCrawlerServiceClient(conn)
		log.Printf("Connected to %s, success!\n", PORT)
		clients = append(clients, client)
	}
	// 分发client
	clientChan := make(chan CrawlerServiceClient, 100)
	go func() {
		for {
			// 轮流分发(轮询)
			for _, client := range clients {
				clientChan <- client
			}
		}
	}()
	return clientChan
}
