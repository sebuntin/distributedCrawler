package doubanparser

import (
	"fetcher"
	"fmt"
	"testing"
	"time"
)

func TestParseBookInfo(t *testing.T) {
	URLList := []string{"https://book.douban.com/subject/1770782/",
		"https://book.douban.com/subject/10554308/",
		"https://book.douban.com/subject/25862578/"}
	Referer := "https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4"
	NameList := []string{
		"追风筝的人",
		"白夜行",
		"解忧杂货店",
	}
	for i, doubanURL := range URLList {
		body, err := fetcher.Fetch(doubanURL, Referer)
		if err != nil {
			t.Fail()
		}
		info := ParseBookInfo(body, NameList[i], doubanURL)
		fmt.Printf("book info: %#v\n", info.Items[0])
		time.Sleep(time.Millisecond * 200)
	}
}
