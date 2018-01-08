package cmd

import (
	"github.com/spf13/cobra"

	"github.com/fujimisakari/go-crawler/backend/server"
)

func serverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start API Server",
		Run:   serve,
	}
	return cmd
}

func serve(cmd *cobra.Command, args []string) {
	server.Serve()
}
