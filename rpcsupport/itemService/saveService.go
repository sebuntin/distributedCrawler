package itemService

import (
	"context"
	"encoding/json"
	"engine"

	"github.com/olivere/elastic/v7"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (this *ItemSaveService) SaveItem(_ context.Context, r *SaveRequest) (*SaveResponse, error) {
	item := engine.Item{}
	err := json.Unmarshal([]byte(r.Item), &item)
	if err != nil {
		return nil, err
	}
	err = elasticSave(item, this.Index, this.Client)
	if err != nil {
		return nil, err
	}
	return &SaveResponse{Status: "ok"}, nil
}

func elasticSave(item engine.Item, index string, esClient *elastic.Client)error {
	// 如果es跑在docker上，需要关掉sniff
	//opt := elastic.SetSniff(false)
	// 存数据
	indexService := esClient.Index().
		Index(index).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}