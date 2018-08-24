package parser

import (
	"golang-demos/crawler-multi-thread/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityURLRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

// ParseCity parser
func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			URL: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	matches = cityURLRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				URL:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}
