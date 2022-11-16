package initialize

import (
	"tsf-cron/config"
	"tsf-cron/pkg/core/log"

	"github.com/elastic/go-elasticsearch/v6"
)

func GetES() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{config.GetString("ES.Address")},
		Username:  config.GetString("ES.Username"),
		Password:  config.GetString("ES.Passwd"),
	}

	c, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Can not connect Elasticsearch %s", err)
	}

	info, err := c.Info()
	if err != nil {
		log.Fatalf("Connect Elasticsearch Err: %s", err)
	}
	if info.StatusCode != 200 {
		log.Fatalf("Connect Elasticsearch Err: %s", info)
	}

	return c
}
