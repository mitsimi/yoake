package cmd

import "github.com/spf13/cobra"

var stopCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(stopCmd)
}
