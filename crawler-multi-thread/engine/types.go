package engine

// Request struct
type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult struct
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// Item interface
type Item struct {
	URL     string
	ID      string
	Payload interface{}
}

// NilParser for default parse logic
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
