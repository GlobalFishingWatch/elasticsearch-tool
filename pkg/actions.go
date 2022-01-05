package pkg

import (
	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/action"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
)

func AddAlias(params types.AddAliasParams) {
	action.AddAlias(params)
}

func CreateIndexWithCustomMapping(params types.CreateIndexParams) {
	action.CreateIndexWithCustomMapping(params)

}

func DeleteIndex(params types.DeleteIndexParams) {
	action.DeleteIndex(params)
}

func DeleteIndexIfExists(params types.DeleteIndexParams) {
	action.DeleteIndexIfExists(params)
}

func DeleteIndicesByPrefix(params types.DeleteIndicesByPrefixParams) {
	action.DeleteIndicesByPrefix(params)
}
