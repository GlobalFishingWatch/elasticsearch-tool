package cmd

import (
	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/action"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	deleteIndexCmd.Flags().StringP( "index-name", "", "", "The name of the destination index (required)")
	deleteIndexCmd.MarkFlagRequired("index-name")
	deleteIndexCmd.Flags().StringP( "elastic-search-url", "", "", "URL exposed by Elasticsearch cluster (required)")
	deleteIndexCmd.MarkFlagRequired("elastic-search-url")

	viper.BindPFlag("index-name", deleteIndexCmd.Flags().Lookup("index-name"))
	viper.BindPFlag("elastic-search-url", deleteIndexCmd.Flags().Lookup("elastic-search-url"))

	rootCmd.AddCommand(deleteIndexCmd)
}

var deleteIndexCmd = &cobra.Command{
	Use:   "delete-index",
	Short: "Delete index",
	Long:  `Delete index
Format:
	elasticsearch-tool delete-index --index-name=[name] --elastic-search-url=[url]
Example:
	elasticsearch-tool delete-index --index-name=test-vessels --elastic-search-url="https://user:password@elastic.gfw.org"`,
	Run: func(cmd *cobra.Command, args []string) {
		params := types.DeleteIndexParams{
			IndexName:        viper.GetString("index-name"),
			ElasticSearchUrl: viper.GetString("elastic-search-url"),
		}
		log.Println("→ Executing Delete Index command")
		action.DeleteIndex(params)
		log.Println("→ Delete Index command finished")
	},
}

