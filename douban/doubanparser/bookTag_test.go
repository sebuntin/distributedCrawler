package doubanparser

import (
	"fetcher"
	"fmt"
	"testing"
)

func TestParseBookTag(t *testing.T) {
	doubanURL := "https://book.douban.com/tag/?view=type&icn=index-sorttags-all"
	body, err := fetcher.Fetch(doubanURL)
	fmt.Println(string(body))
	if err != nil {
		t.Fail()
	}
	doubanPrefix := "https://book.douban.com/"
	parseRet := ParseBookTag(body, doubanPrefix, doubanPrefix)
	for _, item := range parseRet.Items {
		fmt.Printf("Tag: %s\n", item)
	}
}
