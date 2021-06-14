package cmd

import (
	"fmt"
	"os"

	"github.com/misbahulard/shards-ai/config"
	shardsai "github.com/misbahulard/shards-ai/shards-ai"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shards-ai",
	Short: "Shards AI help you to set dynamic number of shards in the elasticsearch indices.",
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
		shardsai.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initialize() {
	config.New()
	config.ConfigureLogger()
	config.ConfigureElasticsearch()
	log.Infof("Starting Shards AI v%s", Version)
}
