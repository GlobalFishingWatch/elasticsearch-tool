package action

import (
	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/common"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
	"github.com/GlobalFishingWatch/elasticsearch-tool/utils"
	"log"
)


func DeleteIndexIfExists(params types.DeleteIndexParams) {
	utils.ValidateUrl(params.ElasticSearchUrl)
	elasticSearchClient = common.CreateElasticSearchClient(params.ElasticSearchUrl)
	deleteIndexIfExists(params.IndexName)
}


func deleteIndexIfExists(indexName string) {
	log.Printf("→ ES →→ Deleting index %s", indexName)
	res, err := elasticSearchClient.Indices.Delete([]string{indexName})
	if err != nil {
		log.Fatalf("→ ES →→ Cannot delete index: %s", err)
	}
	if res.IsError() {
		log.Printf("→ ES →→ Cannot delete index: %s", res)
	}
}