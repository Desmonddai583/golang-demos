package main

import (
	"golang-demos/crawler-multi-thread/engine"
	"golang-demos/crawler-multi-thread/scheduler"
	"golang-demos/crawler-multi-thread/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		URL:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
