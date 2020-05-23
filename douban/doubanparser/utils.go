package doubanparser

import (
	"config"
	"engine"
	"errors"
)

type DoubanParseFuncConversion struct {
}

func (d DoubanParseFuncConversion) Serialize(request engine.Request) (string, []string) {
	return request.Parser.Serialize()
}

func (d DoubanParseFuncConversion) Deserialize(parseFuncName string) (engine.ParseFunc, error) {
	switch parseFuncName {
	case config.ParseBookInfo:
		return ParseBookInfo, nil
	case config.ParseBookList:
		return ParseBookList, nil
	case config.ParseBookTag:
		return ParseBookTag, nil
	}
	return nil, errors.New("Conversion: Valid parseFunc name\n")
}
