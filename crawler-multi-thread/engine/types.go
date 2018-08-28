package engine

// ParserFunc get parse result
type ParserFunc func(contents []byte, url string) ParseResult

// Request struct
type Request struct {
	URL        string
	ParserFunc func(contetns []byte, url string) ParseResult
}

// ParseResult struct
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// Item interface
type Item struct {
	URL     string
	Type    string
	ID      string
	Payload interface{}
}

// NilParser for default parse logic
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
