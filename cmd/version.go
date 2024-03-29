package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is Yoake's`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Yoake Application Autostart Manager v%s\n", rootCmd.Version)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
