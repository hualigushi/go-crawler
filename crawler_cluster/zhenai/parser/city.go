package parser

import (
	"crawler/crawler_distributed/config"
	"crawler/engine"
	"regexp"
)

var (
  profileRe = regexp.MustCompile( `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
  cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity (contents []byte, _ string) engine.ParserResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return result
}
