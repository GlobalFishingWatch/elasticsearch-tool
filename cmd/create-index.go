package cmd

import (
	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/action"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	createIndexCmd.Flags().StringP("mapping", "", "", "The mapping of the destination index (required)")
	createIndexCmd.MarkFlagRequired("mapping")
	createIndexCmd.Flags().StringP("settings", "", "", "The settings of the destination index (optional)")
	createIndexCmd.Flags().StringP("index-name", "", "", "The settings of the destination index (required)")
	createIndexCmd.MarkFlagRequired("index-name")
	createIndexCmd.Flags().StringP("elastic-search-url", "", "", "URL exposed by Elasticsearch cluster (required)")
	createIndexCmd.MarkFlagRequired("elastic-search-url")

	viper.BindPFlag("create-index-mapping", createIndexCmd.Flags().Lookup("mapping"))
	viper.BindPFlag("create-index-settings", createIndexCmd.Flags().Lookup("settings"))
	viper.BindPFlag("create-index-index-name", createIndexCmd.Flags().Lookup("index-name"))
	viper.BindPFlag("create-index-elastic-search-url", createIndexCmd.Flags().Lookup("elastic-search-url"))

	rootCmd.AddCommand(createIndexCmd)
}

var createIndexCmd = &cobra.Command{
	Use:   "create-index",
	Short: "Create new index applying a custom mapping",
	Long: `Create new index applying a custom mapping
Format:
	elasticsearch-tool create-index --mapping=[mapping] --index-name=[name] --elastic-search-url=[url]
Example:
	elasticsearch-tool create-index --mapping={} --index-name=test-vessels --elastic-search-url="https://user:password@elastic.gfw.org"`,
	Run: func(cmd *cobra.Command, args []string) {
		params := types.CreateIndexParams{
			IndexName:        viper.GetString("create-index-index-name"),
			Mapping:          viper.GetString("create-index-mapping"),
			Settings:         viper.GetString("create-index-settings"),
			ElasticSearchUrl: viper.GetString("create-index-elastic-search-url"),
		}
		log.Println("→ Executing Create Index command")
		action.CreateIndexWithCustomMapping(params)
		log.Println("→ Create Index command finished")
	},
}
