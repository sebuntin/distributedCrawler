package main

import (
	. "crawlerService"
	. "doubanparser"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	//端口
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	rpcServer := grpc.NewServer()
	conv := DoubanParseFuncConversion{}
	service := &CrawlerService{
		Conversion: conv,
	}
	RegisterCrawlerServiceServer(rpcServer, service)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	log.Printf("Crawler Service start success ... PORT %d\n", *port)
	err = rpcServer.Serve(conn)
	if err != nil {
		log.Fatalf("Crawler Service: error %s\n", err.Error())
	}
}
