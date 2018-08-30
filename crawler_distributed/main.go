package main

import (
	"fmt"
	"golang-demos/crawler-multi-thread/engine"
	"golang-demos/crawler-multi-thread/scheduler"
	"golang-demos/crawler-multi-thread/zhenai/parser"
	"golang-demos/crawler_distributed/config"
	"golang-demos/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		URL:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// e.Run(engine.Request{
	// 	URL:        "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}
