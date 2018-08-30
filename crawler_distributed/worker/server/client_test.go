package main

import (
	"fmt"
	"golang-demos/crawler_distributed/config"
	"golang-demos/crawler_distributed/rpcsupport"
	"golang-demos/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRPC(
		host,
		worker.CrawlService{},
	)
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		URL: "http://album.zhenai.com/u/108906739",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "安静的雪",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRPC, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
