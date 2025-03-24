package database

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func InitElasticsearch(config struct {
	URL string
}) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{config.URL},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	_, err = es.Info()
	if err != nil {
		return nil, err
	}

	return es, nil
}
