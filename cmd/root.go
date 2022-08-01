package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/mitsimi/yoake/env"
	"github.com/spf13/cobra"
)

var (
	config env.Configuration
)

var rootCmd = &cobra.Command{
	Use:     "yoake",
	Short:   "Little cli program for easy and universal application start on startup",
	Long:    `A simple yet intuitive cli tool for managing windows and linux start up applications.`,
	Version: "1.0.0",
	ValidArgs: []string{
		"setup",
		"enable", "disable",
		"add", "remove",
		"list",
		"delay",
		"help",
		"version",
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !config.Conf.Enabled {
			fmt.Println("yoake is disabled. Use 'yoake enable' to enable it.")
			return
		}

		time.Sleep(time.Duration(config.Conf.Delay) * time.Second)
		for _, app := range config.Apps {
			if !app.Enabled {
				continue
			}
			exec.Command(app.Path).Start()
		}
	},
	DisableFlagsInUseLine: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ExecuteArgs(args []string) {
	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Initialize config directory and file
	var err error
	config, err = env.LoadConfig()
	if err != nil {
		cobra.CheckErr(err)
	}
}
