package model

import "encoding/json"

// douban book
type Book struct {
	Name        string  `json:"Name"`
	Author      string  `json:"Author"`
	Press       string  `json:"Press"`
	OrigName    string  `json:"OrigName"`
	PageNum     int     `json:"PageNum"`
	Price       string  `json:"Price"`
	DoubanScore float64 `json:"DoubanScore"`
	BriefIntro  string  `json:"BriefIntro"`
	AuthorIntro string  `json:"AuthorIntro"`
}

func FromJsonObj(o interface{}) (Book, error) {
	var book Book
	objStr, err := json.Marshal(o)
	if err != nil {
		return book, err
	}
	err = json.Unmarshal(objStr, &book)
	if err != nil {
		return book, err
	}
	return book, nil
}
