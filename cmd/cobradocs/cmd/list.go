package cmd

import (
	"github.com/jrbeverly/cobra-cmd-with-docs/pkg/cobrago"
	"github.com/spf13/cobra"
)

var bucketName string

func init() {
	listS3Cmd.Flags().StringVar(&bucketName, "bucket", "", "name of s3 bucket")
	listS3Cmd.MarkFlagRequired("bucket")

	RootCmd.AddCommand(listS3Cmd)
}

var listS3Cmd = &cobra.Command{
	Use: "list-s3",
	Run: func(cmd *cobra.Command, args []string) {
		cobrago.ListFilesFromStorage(bucketName, storage, writer)
	},
}
