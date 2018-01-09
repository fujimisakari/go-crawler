package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	app "github.com/fujimisakari/go-crawler/backend/application/crawler"
)

func crawlCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "crawl",
		Short: "Start crawl",
		Run:   crawl,
	}
	return cmd
}

func crawl(cmd *cobra.Command, args []string) {
	cmdMapper := app.GetCmdMapper()
	if len(args) == 0 {
		outputCrawlErrorMsg("Crawl target is not specified.")
		return
	}

	crawlTarget := args[0]
	if crawler, ok := cmdMapper[crawlTarget].(app.Crawler); ok {
		crawler.Crawl()
	} else {
		outputCrawlErrorMsg("Crawl target does not exist.")
	}
}

func outputCrawlErrorMsg(msg string) {
	crawlingTargets := strings.Join(app.GetCmdNameList(), ", ")
	fmt.Println(msg)
	fmt.Println(fmt.Sprintf("Target: %s", crawlingTargets))
}
