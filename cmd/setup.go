package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
)

var startUp env.Startup

// TODO: Write descriptions
var setupCmd = &cobra.Command{
	Use:     "setup",
	Aliases: []string{"init"},
	Short:   "Setup starigo autostart",
	Long:    ``,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check operating system
		switch runtime.GOOS {
		case "windows": // Windows
			startUp.Dir = env.Win_StartupDir()
			startUp.File = filepath.Join(startUp.Dir, "starigo.bat")
			startUp.Content = env.Win_Script()

		case "linux": // Linux
			startUp.Dir = env.Linux_StartupDir()
			startUp.File = filepath.Join(startUp.Dir, "starigo.desktop")
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
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
