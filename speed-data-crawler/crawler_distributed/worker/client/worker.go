package client

import (
	"golang-demos/crawler-multi-thread/engine"
	"golang-demos/crawler_distributed/config"
	"golang-demos/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor, error) {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRPC, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
