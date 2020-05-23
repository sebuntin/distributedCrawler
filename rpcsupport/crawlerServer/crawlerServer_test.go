package main

import (
	"config"
	"context"
	. "crawlerService"
	"fmt"
	"log"
	"testing"

	"google.golang.org/grpc"
)

func TestCrawlerService(t *testing.T) {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", config.CrawlerPort), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := NewCrawlerServiceClient(conn)
	//URL := "https://book.douban.com/tag/?view=type&icn=index-sorttags-all"
	URL := "https://book.douban.com/subject/4913064/"
	var request Request
	request.CurUrl = URL
	request.RefUrl = "https://book.douban.com/tag/%E4%B8%AD%E5%9B%BD%E6%96%87%E5%AD%A6"
	request.Parser = &SerializedParser{
		Name: "ParseBookInfo",
		Args: []string{URL, "活着"},
	}
	result, err := client.Process(context.Background(), &request)
	if err != nil {
		log.Fatalf("Crawler Service: process error %v\n", err)
	}
	for i := range result.Items {
		fmt.Println(result.Items[i])
	}

}
