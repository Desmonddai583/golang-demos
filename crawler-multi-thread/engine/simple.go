package engine

import (
	"golang-demos/crawler-multi-thread/fetcher"
	"log"
)

// SimpleEngine struct
type SimpleEngine struct{}

// Run fetch and parse flow
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	// log.Printf("Fetching %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fecher: error"+
			"fetching url %s: %v",
			r.URL, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
