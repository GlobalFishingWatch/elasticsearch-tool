package types

type CreateIndexParams struct {
	IndexName        string
	Mapping          string
	Settings         string
	ElasticSearchUrl string
}
