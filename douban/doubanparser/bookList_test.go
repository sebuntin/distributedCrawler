package doubanparser

import (
	"fetcher"
	"fmt"
	"testing"
)

func TestParseBookList(t *testing.T) {
	doubanURL := "https://book.douban.com/tag/%E5%A4%96%E5%9B%BD%E6%96%87%E5%AD%A6"
	body, err := fetcher.Fetch(doubanURL, "https://www.baidu.com")
	if err != nil {
		t.Fail()
	}
	parseRet := ParseBookList(body, doubanURL)
	for _, item := range parseRet.Items {
		fmt.Printf("Book: %s\n", item)
	}
}