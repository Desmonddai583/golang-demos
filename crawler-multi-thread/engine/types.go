package engine

// ParserFunc get parse result
type ParserFunc func(contents []byte, url string) ParseResult

// Parser interface
type Parser interface {
	Parse(content []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// Request struct
type Request struct {
	URL    string
	Parser Parser
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

// NilParser struct
type NilParser struct{}

// Parse NilParser
func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

// Serialize NilParser
func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// FuncParser struct
type FuncParser struct {
	parser ParserFunc
	name   string
}

// NewFuncParser create func parser
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

// Parse FuncParser
func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

// Serialize FuncParser
func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}
