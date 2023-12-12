package connection

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func EsConn() *elasticsearch.TypedClient {
	cfg := elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTIC_HOST")},
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
		return nil
	}
	return es
}
