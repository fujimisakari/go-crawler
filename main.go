package main

import (
	"fmt"
	"os"

	"github.com/fujimisakari/go-crawler/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
