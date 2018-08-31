package main

import (
	"fmt"
	"golang-demos/crawler_distributed/config"
	"golang-demos/crawler_distributed/rpcsupport"
	"golang-demos/crawler_distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRPC(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{},
	))
}
