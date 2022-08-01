package cmd

import (
	"fmt"

	"github.com/mitsimi/yoake/env"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:     "enable [app]",
	Aliases: []string{"en"},
	Short:   "Enable yoake itself or an application",
	Long: `Enable yoake to start up applications automatically on login or enable an application for the start up.
	Use yoake enable to enable yoake itself.
	Use yoake enable <app> to enable an application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			config.Conf.Enabled = true
			fmt.Printf("yoake is now enabled!")
		} else {
			app := config.Apps[args[0]]
			app.Enabled = true
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
	rootCmd.AddCommand(enableCmd)
}
