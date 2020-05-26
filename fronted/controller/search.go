package controller

import (
	"context"
	"engine"
	"io"
	"model"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"view"

	"github.com/olivere/elastic/v7"
)

type SearchHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func (this *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}
	err = this.getSearchResult(w, q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (this *SearchHandler) getSearchResult(w io.Writer, q string, from int) error {
	search := this.client.Search("douban_book")
	resp, err := search.Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return err
	}
	hits := resp.TotalHits()
	result := resp.Each(reflect.TypeOf(engine.Item{}))
	pageResult := model.SearchResult{}
	pageResult.Items = result
	pageResult.Hits = hits
	pageResult.Start = from
	pageResult.PrevFrom = pageResult.Start - len(result)
	pageResult.NextFrom = pageResult.Start + len(result)
	pageResult.LastFrom = int(pageResult.Hits) - len(result)
	pageResult.Query = q
	err = this.view.Render(w, pageResult)
	if err != nil {
		return err
	}
	return nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Za-z]*):`)
	return re.ReplaceAllString(q, "PayLoad.$1:")
}

func CreatSearchHandler(template string) *SearchHandler {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	return &SearchHandler{
		view.CreateSearchResultView(template),
		client,
	}
}
