package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var globalUsage = `Usage: football [command]
Commands:
  version   Print the version of the football CLI
`

var rootCmd = &cobra.Command{
	Use:  "football",
	Long: globalUsage,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
