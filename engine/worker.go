package engine

import (
	"fetcher"
	"fmt"
	"log"
)

func Worker(r Request) (ParserResult, error) {
	log.Printf("Fecting url: %s\n", r.CurURL)
	body, err := fetcher.Fetch(r.CurURL, r.RefURL)
	if err != nil {
		return ParserResult{}, fmt.Errorf("Fetcher: error "+
			"fetching url %s: %v\n", r.CurURL, err)
	}
	parserResult := r.Parser.Parse(body, r.Args...)
	return parserResult, nil
}
