package cmd

import "github.com/spf13/cobra"

var enableCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(enableCmd)
}
