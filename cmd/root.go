package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd has commands for go-crawler
var RootCmd = &cobra.Command{
	Use:           "go-crawler",
	Short:         "The Tool which is crawling and display",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		crawlCmd(),
	)
}
