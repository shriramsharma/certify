package cmd

import "github.com/spf13/cobra"

var (
	VERSION string
	COMMIT  string
	DATE    string
)

// RootCmd RootCmd
var RootCmd = &cobra.Command{
	Use:   "certify",
	Short: "Verify your certificate and the chain",
	Long:  `Verify your certificate and the chain`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute(version, commit, date string) {
	VERSION = version
	COMMIT = commit
	DATE = date
	RootCmd.Execute()
}
