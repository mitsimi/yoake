package cmd

import "github.com/spf13/cobra"

var startCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(startCmd)
}
