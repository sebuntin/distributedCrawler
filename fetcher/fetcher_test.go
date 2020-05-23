package fetcher

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/parnurzeal/gorequest"
)

func TestFetch(t *testing.T) {
	doubanurl := "https://movie.douban.com/"
	//referer:= "https://movie.douban.com/"
	//body, err := Fetch(doubanurl, referer)
	request := gorequest.New().Proxy("http://118.181.226.166:44640")
	resp, body, err := request.Get(doubanurl).End()
	if err != nil {
		t.Fatalf("Fetcher: fetcher error %s", err[0].Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Fetcher: wrong status code %d", resp.StatusCode)
	}
	fmt.Println(string(body))
}
