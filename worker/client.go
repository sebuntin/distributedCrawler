package worker

import (
	"context"
	. "crawlerService"
	"encoding/json"
	"engine"
	"log"
)

// CRAWLER_PORT

func CreatProcessor(clientChan chan CrawlerServiceClient,
	conv engine.ParseFunConversion) engine.Processor {
	//// grpc dial
	//PORT := fmt.Sprintf(":%d", config.CrawlerPort)
	//conn, err := grpc.Dial(PORT, grpc.WithInsecure())
	//if err != nil {
	//	panic(err)
	//}
	//client := NewCrawlerServiceClient(conn)
	return func(request engine.Request) (engine.ParserResult, error) {
		// Serialize engine.Request
		var result engine.ParserResult
		var r Request
		r.RefUrl = request.RefURL
		r.CurUrl = request.CurURL
		parseFunName, args := request.Parser.Serialize()
		r.Parser = &SerializedParser{
			Name: parseFunName,
			Args: args,
		}
		client := <-clientChan
		resp, err := client.Process(context.Background(), &r)
		if err != nil {
			return result, err
		}
		for i := range resp.Items {
			var item engine.Item
			err := json.Unmarshal([]byte(resp.Items[i]), &item)
			if err != nil {
				log.Printf("Crawler Service: marshal "+
					"item %v error %v", resp.Items[i], err)
				continue
			}
			result.Items = append(result.Items, item)
		}
		for i := range resp.Requests {
			var req engine.Request
			req.Args = resp.Requests[i].Parser.Args
			req.CurURL = resp.Requests[i].CurUrl
			req.RefURL = resp.Requests[i].RefUrl
			parseFunc, err := conv.Deserialize(resp.Requests[i].Parser.Name)
			if err != nil {
				log.Printf("Crawler Service: "+
					"error %v\n", err)
				continue
			}
			req.Parser = engine.NewFuncParser(resp.Requests[i].Parser.Name, req.Args, parseFunc)
			result.Requests = append(result.Requests, req)
		}
		return result, nil
	}
}
