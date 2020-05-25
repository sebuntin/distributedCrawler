package persist

import (
	"context"
	"encoding/json"
	"engine"
	"fmt"
	. "itemService"
	"log"

	"google.golang.org/grpc"
)

func ItemSaver(host int) chan engine.Item {
	//client, err := elastic.NewClient()
	conn, err := grpc.Dial(fmt.Sprintf(":%d", host), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// get rpc client
	client := NewSaveServiceClient(conn)
	out := make(chan engine.Item)
	go func() {
		count := 0
		for {
			item := <-out
			count++
			log.Printf("Item Saver: got item #%d, %v", count, item)
			if item.PayLoad != nil {
				// call rpc to save item
				itemStr, err := json.Marshal(item)
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
