package crawlerService

import (
	"context"
	"encoding/json"
	"engine"
	"log"
)

type CrawlerService struct {
	Conversion engine.ParseFunConversion
}

func (this *CrawlerService) Process(_ context.Context, r *Request) (*Response, error) {
	request, err := deserializeRequest(*r, this.Conversion)
	if err != nil {
		return nil, err
	}
	parseResult, err := engine.Worker(request)
	if err != nil {
		return nil, err
	}
	var resp Response
	serializeResponse(parseResult, &resp, this.Conversion)
	return &resp, nil
}

// 将request反序列化为engine.Request
func deserializeRequest(r Request, conv engine.ParseFunConversion) (engine.Request, error) {
	var request engine.Request
	request.CurURL = r.CurUrl
	request.RefURL = r.RefUrl
	request.Args = r.Parser.Args
	parseFunc, err := conv.Deserialize(r.Parser.Name)
	if err != nil {
		return request, err
	}
	request.Parser = engine.NewFuncParser(r.Parser.Name, r.Parser.Args, parseFunc)
	return request, nil
}

// 将result序列化为Response
func serializeResponse(result engine.ParserResult, response *Response, conv engine.ParseFunConversion) {
	for i := range result.Items {
		// serialize items
		marshal, err := json.Marshal(result.Items[i])
		if err != nil {
			log.Printf("Crawler Service: error item %v\n", result.Items[i])
			continue
		}
		response.Items = append(response.Items, string(marshal))
	}
	// serialize requests
	for i := range result.Requests {
		var request Request
		var serializedParser SerializedParser
		request.RefUrl = result.Requests[i].RefURL
		request.CurUrl = result.Requests[i].CurURL
		parseFuncName, args := conv.Serialize(result.Requests[i])
		serializedParser.Name = parseFuncName
		serializedParser.Args = args
		request.Parser = &serializedParser
		response.Requests = append(response.Requests, &request)
	}
}
