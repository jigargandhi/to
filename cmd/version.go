package cmd

import (
	"fmt"

	"github.com/jigargandhi/to/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "prints version information",
	Long:  `prints version information of the tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("to version %s commit %s\n", version.Version(), version.Commit())
	},
}
