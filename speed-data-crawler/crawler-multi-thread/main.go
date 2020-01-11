package main

import (
	"golang-demos/crawler-multi-thread/config"
	"golang-demos/crawler-multi-thread/engine"
	"golang-demos/crawler-multi-thread/persist"
	"golang-demos/crawler-multi-thread/scheduler"
	"golang-demos/crawler-multi-thread/xcar/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		URL:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCarList, config.ParseCarList),
	})

	// e.Run(engine.Request{
	// 	URL:        "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}
