package action

import (
	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/common"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
	"github.com/GlobalFishingWatch/elasticsearch-tool/utils"
	"log"
)

func AddAlias(params types.AddAliasParams) {
	utils.ValidateUrl(params.ElasticSearchUrl)
	elasticSearchClient = common.CreateElasticSearchClient(params.ElasticSearchUrl)
	createAlias(params.IndexName, params.Alias)
}

func createAlias(indexName string, alias string) {
	indices := []string{indexName}
	res, err := elasticSearchClient.Indices.PutAlias(indices, alias)
	if err != nil {
		log.Fatalf("→ ES →→ Error creating new alias: %s", err)
	}
	log.Printf("→ ES →→ Create Alias response: %v", res)
}