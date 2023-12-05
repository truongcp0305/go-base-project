package connection

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func EsConn() *elasticsearch.TypedClient {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
		return nil
	}
	// _, err = es.Indices.Create(os.Getenv("INDEX")).Do(context.Background())
	// if err != nil {
	// 	log.Fatalf("Error creating index: %s", err)
	// 	return nil
	// }
	return es
}
