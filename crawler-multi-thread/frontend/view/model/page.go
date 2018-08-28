package model

import "golang-demos/crawler-multi-thread/engine"

// SearchResult struct
type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
