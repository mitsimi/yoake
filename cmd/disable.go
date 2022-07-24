package cmd

import (
	"fmt"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:     "disable [app]",
	Aliases: []string{"dis"},
	Short:   "Disable starigo itself or an application",
	Long: `Disable starigo to start up applications automatically on login or disable an application for the start up.
	Use "starigo disable" to enable starigo itself.
	Use "starigo disable <app>" to Disable an application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			config.Conf.Enabled = false
			fmt.Printf("starigo is now disabled!")
		} else {
			app := config.Apps[args[0]]
			app.Enabled = false
			config.Apps[args[0]] = app

			fmt.Printf("%s is now disabled!", args[0])
		}

		if err := config.WriteConfig(); err != nil {
			env.WriteLog(err.Error())
			cobra.CheckErr(err)
		}
	}}

func init() {
	rootCmd.AddCommand(disableCmd)
}
