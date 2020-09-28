package cmd

import (
	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/action"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	addAliasCmd.Flags().StringP("index-name", "i", "", "The name of the index to create alias")
	addAliasCmd.MarkFlagRequired("index-name")

	addAliasCmd.Flags().StringP("alias", "a", "", "Alias name")
	addAliasCmd.MarkFlagRequired("alias")

	addAliasCmd.Flags().StringP( "elastic-search-url", "u", "", "URL exposed by Elasticsearch cluster (required)")
	addAliasCmd.MarkFlagRequired("elastic-search-url")

	viper.BindPFlag("add-alias-index-name", addAliasCmd.Flags().Lookup("index-name"))
	viper.BindPFlag("add-alias-alias", addAliasCmd.Flags().Lookup("alias"))
	viper.BindPFlag("add-alias-elastic-search-url", addAliasCmd.Flags().Lookup("elastic-search-url"))

	rootCmd.AddCommand(addAliasCmd)
}

var addAliasCmd = &cobra.Command{
	Use:   "add-alias",
	Short: "Add an alias to an index",
	Long:  `Adds an alias to an index
Format:
	bq2es add-alias --index-name=[name] --alias=[name] --elastic-search-url=[url]
Example:
	bq2es add-alias --index-name=gfw-tasks-2020 --alias=gfw-tasks --elastic-search-url=https://user:password@elastic.gfw.org`,
	Run: func(cmd *cobra.Command, args []string) {
		params := types.AddAliasParams{
			IndexName: viper.GetString("add-alias-index-name"),
			Alias:                viper.GetString("add-alias-alias"),
			ElasticSearchUrl: viper.GetString("add-alias-elastic-search-url"),
		}

		log.Println("→ Executing Add Alias command")
		action.AddAlias(params)
		log.Println("→ Execution completed")
	},
}

