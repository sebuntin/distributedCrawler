package main

import (
	. "config"
	"flag"
	"fmt"
	. "itemService"
	"log"
	"net"

	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc"
)

//const ITEM_PORT = ":1234"
var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	//端口
	flag.Parse()
	//端口
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	rpcServer := grpc.NewServer()
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	log.Printf("elastic search client start success ... ")
	service := &ItemSaveService{
		Client: client,
		Index:  ElasticIndex,
	}
	RegisterSaveServiceServer(rpcServer, service)

	conn, err := net.Listen("tcp",
		fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	log.Printf("Item Service start success ... PORT :%d\n", *port)
	err = rpcServer.Serve(conn)
	if err != nil {
		log.Fatalf("Item Service: error %s\n", err.Error())
	}
}
