package cmd

import (
	"github.com/jigargandhi/to/url"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

var versionCmd = &cobra.Command{
	Use:   "url",
	Short: "to url tagName parameter",
	Long:  `Opens a url configured in tagName with parameter replacing the TAG`,
	Run: func(cmd *cobra.Command, args []string) {
		url.Go(args[0], args[1])
	},
}
