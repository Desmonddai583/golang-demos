package config

const (
	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// Service ports
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	// ElasticSearch
	ElasticIndex = "dating_profile"

	// RPC Endpoints
	ItemSaverRPC    = "ItemSaverService.Save"
	CrawlServiceRPC = "CrawlService.Process"
)
