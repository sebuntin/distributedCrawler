package persist

import (
	"config"
	"context"
	"encoding/json"
	"engine"
	"fmt"
	. "itemService"
	"log"

	"google.golang.org/grpc"
)

func ItemSaver() chan interface{} {
	//client, err := elastic.NewClient()
	conn, err := grpc.Dial(fmt.Sprintf(":%d", config.ItemSaverPort), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// get rpc client
	client := NewSaveServiceClient(conn)
	out := make(chan interface{})
	go func() {
		count := 0
		for {
			item := <-out
			count++
			log.Printf("Item Saver: got item #%d, %v", count, item)
			if book, ok := item.(engine.Item); ok {
				// call rpc to save item
				itemStr, err := json.Marshal(book)
				if err != nil {
					log.Printf("Item Saver: error saving item #%v, %v\n", item, err)
				}
				if _, err := client.SaveItem(context.Background(),
					&SaveRequest{Item: string(itemStr)}); err != nil {
					log.Printf("Item Saver: error saving item %v: %v\n", item, err)
					continue
				}
			}
		}
	}()
	return out
}
