package cmd

import (
	"fmt"
	"strings"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove <app_name>",
	Aliases: []string{"rm"},
	Short:   "Remove an app from your startup list",
	Long: `Remove an app from your startup list
	For example: starigo remove spotify`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.ToLower(args[0])

		if _, ok := config.Apps[name]; !ok {
			fmt.Printf("App %s doesn't exists.\n", name)
			return
		}

		delete(config.Apps, name)

		if err := config.WriteConfig(); err != nil {
			env.WriteLog(err.Error())
			cobra.CheckErr(err)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
