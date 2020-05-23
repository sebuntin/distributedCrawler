package main

import (
	"config"
	. "crawlerService"
	. "doubanparser"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//端口
	rpcServer := grpc.NewServer()
	conv := DoubanParseFuncConversion{}
	service := &CrawlerService{
		Conversion: conv,
	}
	RegisterCrawlerServiceServer(rpcServer, service)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", config.CrawlerPort))
	if err != nil {
		panic(err)
	}
	log.Printf("Crawler Service start success ... PORT %d\n", config.CrawlerPort)
	err = rpcServer.Serve(conn)
	if err != nil {
		log.Fatalf("Crawler Service: error %s\n", err.Error())
	}
}
