package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of certify",
	Long:  `All software has versions. This is certify's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nCommmit: %s\nDate: %s\n", VERSION, COMMIT, DATE)
	},
}
