package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
)

var (
	enabled bool

	startup_dir     string
	startup_file    string
	startup_content func() string
)

var rootCmd = &cobra.Command{
	Use:     "starigo",
	Short:   "Little cli program for easy and universal application start on startup",
	Long:    `A simple yet intuitive cli tool for managing windows and linux start up applications.`,
	Version: "0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		if enabled {
			fmt.Println("enabled")
		} else {
			fmt.Println("disabled")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// Check operating system
	switch runtime.GOOS {
	case "windows": // Windows
		startup_dir = env.WinStartupDir()
		startup_file = startup_dir + "\\starigo.bat"
		startup_content = env.Win_Script

	case "linux": // Linux
		startup_dir = env.LinuxStartupDir()
		startup_file = startup_dir + "/starigo.desktop"
		startup_content = env.Linux_Desktop

	default: // Other (Darwin, FreeBSD, OpenBSD, Plan9, etc.)
		cobra.CheckErr(fmt.Errorf("unsupported OS: %s", runtime.GOOS))
	}

	// Initialize startup directory and file
	go initStartup()
}

func initStartup() {
	// Check if startup directory exists
	if _, err := os.Stat(startup_dir); os.IsNotExist(err) {
		// Create startup directory
		err := os.Mkdir(startup_dir, 0755)
		if err != nil {
			env.WriteLog(err.Error())
		}
	}

	// Check if startup file exists
	if _, err := os.Stat(startup_file); os.IsNotExist(err) {
		// Write startup file
		err := ioutil.WriteFile(startup_file, []byte(startup_content()), 0644)
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

	// Check if config file exists
	if _, err := os.Stat(env.ConfigFile()); os.IsNotExist(err) {
		// TODO: Write config file with viper

	}
}
