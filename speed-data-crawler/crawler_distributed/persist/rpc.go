package persist

import (
	"golang-demos/crawler-multi-thread/engine"
	"golang-demos/crawler-multi-thread/persist"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

// ItemSaverService struct
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// Save save service to save item
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
