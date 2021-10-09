package cmd

import (
	"github.com/jigargandhi/to/url"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(urlCommand)
}

var urlCommand = &cobra.Command{
	Use:   "url",
	Short: "to url tagName [parameter], opens url configured with tagName",
	Long:  `Opens a url configured in tagName with parameter replacing the TAG`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tag := ""
		if len(args) == 2 {
			tag = args[1]
		}
		url.Go(args[0], tag)
	},
}
