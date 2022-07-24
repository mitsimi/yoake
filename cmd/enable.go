package cmd

import (
	"fmt"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:     "enable [app]",
	Aliases: []string{"en"},
	Short:   "Enable starigo itself or an application",
	Long: `Enable starigo to start up applications automatically on login or enable an application for the start up.
	Use starigo enable to enable starigo itself.
	Use starigo enable <app> to enable an application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			config.Conf.Enabled = true
			fmt.Printf("starigo is now enabled!")
		} else {
			app := config.Apps[args[0]]
			app.Enabled = true
			config.Apps[args[0]] = app

			fmt.Printf("%s is now enabled!", args[0])
		}
		if err := config.WriteConfig(); err != nil {
			env.WriteLog(err.Error())
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
