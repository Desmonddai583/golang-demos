package main

import (
	"fmt"
	"golang-demos/crawler-multi-thread/engine"
	"golang-demos/crawler-multi-thread/scheduler"
	"golang-demos/crawler-multi-thread/zhenai/parser"
	"golang-demos/crawler_distributed/config"
	itemsaver "golang-demos/crawler_distributed/persist/client"
	worker "golang-demos/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		URL:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

	// e.Run(engine.Request{
	// 	URL:        "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}
