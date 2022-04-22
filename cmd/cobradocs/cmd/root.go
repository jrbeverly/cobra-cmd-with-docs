package cmd

import (
	"fmt"
	"os"

	"github.com/jrbeverly/cobra-cmd-with-docs/internal/aws"
	"github.com/jrbeverly/cobra-cmd-with-docs/internal/console"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "app",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var storage aws.S3RemoteStorage
var writer console.ConsoleWriter

func Execute() {
	storage = aws.NewS3RemoteStorage()
	writer = console.NewConsoleWriter()
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
