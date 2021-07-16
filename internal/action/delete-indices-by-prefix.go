package action

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/common"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
	"github.com/GlobalFishingWatch/elasticsearch-tool/utils"
)

type Index struct {
	Index string `json:"index"`
}

func DeleteIndicesByPrefix(params types.DeleteIndicesByPrefixParams) {
	utils.ValidateUrl(params.ElasticSearchUrl)
	elasticSearchClient = common.CreateElasticSearchClient(params.ElasticSearchUrl)
	deleteIndicesByPrefix(params.Prefix, params.NoDeleteIndex)
}

func deleteIndicesByPrefix(prefix, noDeleteIndex string) {
	log.Printf("→ ES →→ Listing indices by prefix %s", prefix)
	res, err := elasticSearchClient.Cat.Indices(elasticSearchClient.Cat.Indices.WithIndex(fmt.Sprintf("%s*", prefix)), elasticSearchClient.Cat.Indices.WithFormat("json"))

	if err != nil {
		log.Fatalf("→ ES →→ Cannot list indices with prefix: %s", err)
	}
	if res.IsError() {
		log.Fatalf("→ ES →→ Cannot list indices with prefix: %s", res)
	}
	defer res.Body.Close()
	var indices []Index
	err = json.NewDecoder(res.Body).Decode(&indices)
	if err != nil {
		log.Fatalf("→ ES →→ Cannot list indices with prefix: %s", err)
	}
	for _, index := range indices {
		if index.Index != noDeleteIndex {
			deleteIndex(index.Index)
		} else {
			log.Printf("→ ES →→ %s index not delete because it is in no-delete-index param", index.Index)
		}
	}
}
