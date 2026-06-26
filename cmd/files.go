package cmd

import (
	"github.com/spf13/cobra"
)

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Operations pertaining to files.",
}

func init() {
	rootCmd.AddCommand(filesCmd)
}
