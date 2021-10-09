package cmd

import (
	"github.com/jigargandhi/to/url"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add or replace a shortcut",
	Long:  `Add or replace a shortcut in the shortcuts' file`,
	Run: func(cmd *cobra.Command, args []string) {
		url.Add(args[0], args[1])
	},
}
