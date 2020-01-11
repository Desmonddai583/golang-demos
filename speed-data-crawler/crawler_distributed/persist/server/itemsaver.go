package main

import (
	"flag"
	"fmt"
	"golang-demos/crawler_distributed/config"
	"golang-demos/crawler_distributed/persist"
	"golang-demos/crawler_distributed/rpcsupport"
	"log"

	"github.com/olivere/elastic/v7"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	log.Fatal(serveRPC(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRPC(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRPC(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
