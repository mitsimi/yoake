package cmd

import "github.com/spf13/cobra"

var disableCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(disableCmd)
}
