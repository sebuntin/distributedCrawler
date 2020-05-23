package model

import "encoding/json"

// douban book
type Book struct {
	Name        string
	Author      string
	Press       string
	OrigName    string
	PageNum     int
	Price       string
	DoubanScore float64
	BriefIntro  string
	AuthorIntro string
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
