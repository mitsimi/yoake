package cmd

import (
	"fmt"

	"github.com/mitsimi/yoake/env"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:     "disable [app]",
	Aliases: []string{"dis"},
	Short:   "Disable yoake itself or an application",
	Long: `Disable yoake to start up applications automatically on login or disable an application for the start up.
	Use "yoake disable" to enable yoake itself.
	Use "yoake disable <app>" to Disable an application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			config.Conf.Enabled = false
			fmt.Printf("yoake is now disabled!")
		} else {
			app := config.Apps[args[0]]
			app.Enabled = false
			config.Apps[args[0]] = app

			showCmd.Run(cmd, args)
		}

		if err := config.WriteConfig(); err != nil {
			env.WriteLog(err.Error())
			cobra.CheckErr(err)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
