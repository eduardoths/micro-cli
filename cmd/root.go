package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var buildVersion string

const VERSION_FLAG = "version"

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "microcli",
		Long:    "Microservices tool",
		Version: buildVersion,
		Run:     defaultCommand,
	}
	return cmd
}

func defaultCommand(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Fatalf("Something unexpected happened %s", err.Error())
	}
}

func Execute() {
	root := newRootCommand()
	if err := root.Execute(); err != nil {
		log.Fatalf("Failed to init cli %s", err.Error())
	}
}
