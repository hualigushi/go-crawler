package worker

import (
	"crawler/crawler_distributed/config"
	"crawler/engine"
	"crawler/zhenai/parser"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items []engine.Item
	Requests []Request
}

func SerializeRequest (r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult (r engine.ParserResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest((req)))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url: r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult (r ParseResult) engine.ParserResult {
	result := engine.ParserResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch  p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParserCityList,config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity,config.ParseCity), nil
	case config.NilParser:
		return  engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return  nil, fmt.Errorf("invalid arg: %v", p.Args)
		}

	default:
		return nil, errors.New("unknown parse name")
	}
}

// {"ParseCityList", nil}
// {"ProfileParser", userName}

