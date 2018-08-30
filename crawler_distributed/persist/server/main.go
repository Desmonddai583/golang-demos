package main

import (
	"fmt"
	"golang-demos/crawler_distributed/config"
	"golang-demos/crawler_distributed/persist"
	"golang-demos/crawler_distributed/rpcsupport"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

func main() {
	log.Fatal(serveRPC(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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
