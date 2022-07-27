package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Removing config...")
		err := os.RemoveAll(env.ConfigDir())
		cobra.CheckErr(err)
		fmt.Println("\033[36mDone\033[0m")

		fmt.Print("Removing binary...")
		switch runtime.GOOS {
		case "windows":
			uninstallWindows()
		case "linux":
			uninstallLinux()
		default:
			cobra.CheckErr(fmt.Errorf("unsupported OS: %s", runtime.GOOS))
		}
		fmt.Println("\033[36mDone\033[0m")
	},
	DisableFlagsInUseLine: true,
}

func uninstallWindows() {
	err := os.Remove(filepath.Join(env.Win_StartupDir(), "starigo.bat"))
	if err != nil {
		env.WriteLog(err.Error())
		cobra.CheckErr(err)
	}

	err = exec.Command("cmd", "/C", "del", filepath.Join(env.BinaryDir(), "starigo.exe")).Start()
	if err != nil {
		env.WriteLog(err.Error())
		cobra.CheckErr(err)
	}
}

func uninstallLinux() {
	err := os.Remove(filepath.Join(env.Linux_StartupDir(), "starigo.desktop"))
	if err != nil {
		env.WriteLog(err.Error())
		cobra.CheckErr(err)
	}

	err = exec.Command("rm", filepath.Join(env.BinaryDir(), "starigo")).Start()
	if err != nil {
		env.WriteLog(err.Error())
		cobra.CheckErr(err)
	}
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
