package cmd

import "github.com/spf13/cobra"

var removeCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(removeCmd)
}
