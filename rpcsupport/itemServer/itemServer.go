package main

import (
	. "config"
	. "itemService"
	"log"
	"net"

	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc"
)

const ITEM_PORT = ":1234"

func main() {
	//端口
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

	conn, err := net.Listen("tcp", ITEM_PORT)
	if err != nil {
		panic(err)
	}
	log.Printf("Item Service start success ... PORT %s\n", ITEM_PORT)
	err = rpcServer.Serve(conn)
	if err != nil {
		log.Fatalf("Item Service: error %s\n", err.Error())
	}
}
