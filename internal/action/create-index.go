package action

import (
	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/common"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
	"github.com/GlobalFishingWatch/elasticsearch-tool/utils"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"log"
	"strings"
)

func CreateIndexWithCustomMapping(params types.CreateIndexParams) {
	utils.ValidateUrl(params.ElasticSearchUrl)
	elasticClient := common.CreateElasticSearchClient(params.ElasticSearchUrl)
	createIndex(elasticClient, params.IndexName)
	mapping := parseToReader(params.Mapping)
	res := putMapping(elasticClient, params.IndexName, mapping)
	log.Printf("→ Set Mapping response: %v", res)

	if params.Settings != "" {
		settings := parseToReader(params.Settings)
		res = putSettings(elasticClient, params.IndexName, settings)
		log.Printf("→ Set Settings response: %v", res)
	}
}

func createIndex(elasticSearchClient *elasticsearch.Client, indexName string) {
	log.Printf("→ ES →→ Creating index %s", indexName)
	res, err := elasticSearchClient.Indices.Create(indexName)
	if err != nil {
		log.Fatalf("→ ES →→ Cannot create index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("→ ES →→ Cannot create index: %s", res)
	}
}

func putMapping(elasticClient *elasticsearch.Client, indexName string, mapping io.Reader) *esapi.Response {
	res, err := elasticClient.Indices.PutMapping(mapping, func(index *esapi.IndicesPutMappingRequest) {
		index.Index = []string{indexName}
	})
	if err != nil {
		log.Fatalf("→ ES →→ Cannot put mapping: %s", err)
	}
	if res.IsError() {
		log.Fatalf("→ ES →→ Cannot delete index: %s", res)
	}
	return res
}

func putSettings(elasticClient *elasticsearch.Client, indexName string, settings io.Reader) *esapi.Response {
	res, err := elasticClient.Indices.PutSettings(settings, func(index *esapi.IndicesPutSettingsRequest) {
		index.Index = []string{indexName}
	})
	if err != nil {
		log.Fatalf("→ ES →→ Cannot put mapping: %s", err)
	}
	if res.IsError() {
		log.Fatalf("→ ES →→ Cannot delete index: %s", res)
	}
	return res
}

func parseToReader(mapping string) io.Reader {
	log.Printf("→ ES →→ Casting to reader: %s", mapping)
	return strings.NewReader(mapping)
}
