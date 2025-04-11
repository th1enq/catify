package elastic

import "github.com/elastic/go-elasticsearch"

type Client struct {
	*elasticsearch.Client
}

func Init() (*Client, error) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}
