package cmd

import (
	"log"

	"github.com/GlobalFishingWatch/elasticsearch-tool/internal/action"
	"github.com/GlobalFishingWatch/elasticsearch-tool/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	deleteIndicesByPrefixCmd.Flags().StringP("index-prefix", "", "", "The prefix of the indices to delete (required)")
	deleteIndicesByPrefixCmd.MarkFlagRequired("index-prefix")
	deleteIndicesByPrefixCmd.Flags().StringP("no-delete-index", "", "", "Index name that you do not want to delete (optional)")
	deleteIndicesByPrefixCmd.Flags().StringP("elastic-search-url", "", "", "URL exposed by Elasticsearch cluster (required)")
	deleteIndicesByPrefixCmd.MarkFlagRequired("elastic-search-url")

	viper.BindPFlag("delete-indices-by-prefix-index-prefix", deleteIndicesByPrefixCmd.Flags().Lookup("index-prefix"))
	viper.BindPFlag("delete-indices-by-prefix-no-delete-index", deleteIndicesByPrefixCmd.Flags().Lookup("no-delete-index"))
	viper.BindPFlag("delete-indices-by-prefix-elastic-search-url", deleteIndicesByPrefixCmd.Flags().Lookup("elastic-search-url"))

	rootCmd.AddCommand(deleteIndicesByPrefixCmd)
}

var deleteIndicesByPrefixCmd = &cobra.Command{
	Use:   "delete-indices-by-prefix",
	Short: "Delete indices by prefix",
	Long: `Delete indices by prefix
Format:
	elasticsearch-tool delete-indices-by-prefix --index-prefix=[name] --elastic-search-url=[url]
Example:
	elasticsearch-tool delete-indices --index-prefix=test-vessels --no-delete-index=test-vessels-2021-01 --elastic-search-url="https://user:password@elastic.gfw.org"`,
	Run: func(cmd *cobra.Command, args []string) {
		params := types.DeleteIndicesByPrefixParams{
			Prefix:           viper.GetString("delete-indices-by-prefix-index-prefix"),
			NoDeleteIndex:    viper.GetString("delete-indices-by-prefix-no-delete-index"),
			ElasticSearchUrl: viper.GetString("delete-indices-by-prefix-elastic-search-url"),
		}
		log.Println("→ Executing Delete Indices by prefix command")
		action.DeleteIndicesByPrefix(params)
		log.Println("→ Delete Indices by prefix command finished")
	},
}
