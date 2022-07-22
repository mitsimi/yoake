package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	enabled bool

	startUp env.Startup
)

var rootCmd = &cobra.Command{
	Use:       "starigo",
	Short:     "Little cli program for easy and universal application start on startup",
	Long:      `A simple yet intuitive cli tool for managing windows and linux start up applications.`,
	Version:   "0.1.0",
	ValidArgs: []string{"start", "stop", "enable", "disable", "add", "remove", "list", "delay", "help", "version"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Initialize startup directory and file
	go initStartup()

	// Initialize config directory and file
	go initConfig()
}

func initStartup() {
	// Check operating system
	switch runtime.GOOS {
	case "windows": // Windows
		startUp.Dir = env.Win_StartupDir()
		startUp.File = startUp.Dir + "\\starigo.bat"
		startUp.Content = env.Win_Script()

	case "linux": // Linux
		startUp.Dir = env.Linux_StartupDir()
		startUp.File = startUp.Dir + "/starigo.desktop"
		startUp.Content = env.Linux_Desktop()

	default: // Other (Darwin, FreeBSD, OpenBSD, Plan9, etc.)
		cobra.CheckErr(fmt.Errorf("unsupported OS: %s", runtime.GOOS))
	}

	// Check if startup directory exists
	if _, err := os.Stat(startUp.Dir); os.IsNotExist(err) {
		// Create startup directory
		err := os.Mkdir(startUp.Dir, 0755)
		if err != nil {
			env.WriteLog(err.Error())
		}
	}

	// Check if startup file exists
	if _, err := os.Stat(startUp.File); os.IsNotExist(err) {
		// Write startup file
		err := ioutil.WriteFile(startUp.File, []byte(startUp.Content), 0644)
		if err != nil {
			env.WriteLog(err.Error())
		}
	}
}

// TODO: Initialize config file loading and management
func initConfig() {
	// Check if config directory exists
	if _, err := os.Stat(env.ConfigDir()); os.IsNotExist(err) {
		// Create config directory
		err := os.Mkdir(env.ConfigDir(), 0755)
		if err != nil {
			env.WriteLog(err.Error())
		}
	}

	// Configure viper
	config := viper.New()
	config.AddConfigPath(env.ConfigDir())
	config.SetConfigType("json")
	config.SetConfigName("config")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			env.WriteLog("Config file not found")
			env.WriteLog("Creating new config file")

			// Write config file
			err := config.WriteConfig()
			if err != nil {
				env.WriteLog(err.Error())
			}

		} else {
			// Config file was found but another error was produced
			env.WriteLog(err.Error())
			cobra.CheckErr(err.Error())
		}
	}

	// Config file found and successfully parsed

}
