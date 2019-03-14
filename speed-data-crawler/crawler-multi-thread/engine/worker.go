package engine

import (
	"golang-demos/crawler-multi-thread/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	// log.Printf("Fetching %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fecher: error"+
			"fetching url %s: %v",
			r.URL, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.URL), nil
}
