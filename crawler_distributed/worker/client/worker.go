package client

import (
	"fmt"
	"golang-demos/crawler-multi-thread/engine"
	"golang-demos/crawler_distributed/config"
	"golang-demos/crawler_distributed/rpcsupport"
	"golang-demos/crawler_distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRPC, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
