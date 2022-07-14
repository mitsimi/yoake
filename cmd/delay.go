package cmd

import "github.com/spf13/cobra"

var delayCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(delayCmd)
}
