package main

import (
	"flag"
	"fmt"
	"golang-demos/crawler_distributed/rpcsupport"
	"golang-demos/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	log.Fatal(rpcsupport.ServeRPC(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{},
	))
}
