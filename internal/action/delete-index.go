package action

import (
	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/common"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
	"github.com/GlobalFishingWatch/elasticsearch-tool/utils"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var elasticSearchClient *elasticsearch.Client

func DeleteIndex(params types.DeleteIndexParams) {
	utils.ValidateUrl(params.ElasticSearchUrl)
	elasticSearchClient = common.CreateElasticSearchClient(params.ElasticSearchUrl)
	deleteIndex(params.IndexName)
}


func deleteIndex(indexName string) {
	log.Printf("→ ES →→ Deleting index %s", indexName)
	res, err := elasticSearchClient.Indices.Delete([]string{indexName})
	if err != nil {
		log.Fatalf("→ ES →→ Cannot delete index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("→ ES →→ Cannot delete index: %s", res)
	}
	log.Printf("→ Set Mapping response: %v", res)
}