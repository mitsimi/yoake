package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
)

var (
	config  env.Configuration
	startUp env.Startup
)

var rootCmd = &cobra.Command{
	Use:     "starigo",
	Short:   "Little cli program for easy and universal application start on startup",
	Long:    `A simple yet intuitive cli tool for managing windows and linux start up applications.`,
	Version: "1.0.0",
	ValidArgs: []string{
		"enable", "disable",
		"add", "remove",
		"list",
		"delay",
		"help",
		"version",
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !config.Conf.Enabled {
			fmt.Println("StariGo is disabled. Use 'starigo enable' to enable it.")
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
	var err error
	config, err = env.LoadConfig()
	if err != nil {
		cobra.CheckErr(err)
	}
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
