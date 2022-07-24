package cmd

import (
	"fmt"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <app_name> <app_executable_path>",
	Short: "Add a new app to your startup list",
	Long: `Add a new app to your startup list
	For example: starigo add spotify %USERPROFILE%/Desktop/Spotify.exe`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		path := args[1]

		if _, ok := config.Apps[name]; ok {
			fmt.Printf("App %s already exists.\n", name)
			return
		}

		config.Apps[name] = env.App{
			Enabled: true,
			Path:    path,
		}

		if err := config.WriteConfig(); err != nil {
			env.WriteLog(err.Error())
			cobra.CheckErr(err)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		showCmd.Run(cmd, []string{})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
