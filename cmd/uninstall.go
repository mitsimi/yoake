package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mitsimi/yoake/env"
	"github.com/spf13/cobra"
)

var confirm string

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls the application",
	Long:  `Removes all files and directories created by the application. No recovery is possible.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt for confirmation
		fmt.Print("\033[31mAll files will be deleted permanently. Are you sure you want to uninstall yoake? (y/N) \033[0m")
		fmt.Scanln(&confirm)
		confirm = strings.ToLower(confirm)
		if confirm == "" {
			confirm = "n"
		}

		if confirm[:1] != "y" {
			fmt.Println("\n\033[31mAborted.\033[0m")
			return
		}

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
	err := os.Remove(filepath.Join(env.Win_StartupDir(), "yoake.bat"))
	if err != nil {
		env.WriteLog(err.Error())
		cobra.CheckErr(err)
	}

	err = exec.Command("cmd", "/C", "del", filepath.Join(env.BinaryDir(), "yoake.exe")).Start()
	if err != nil {
		env.WriteLog(err.Error())
		cobra.CheckErr(err)
	}
}

func uninstallLinux() {
	err := os.Remove(filepath.Join(env.Linux_StartupDir(), "yoake.desktop"))
	if err != nil {
		env.WriteLog(err.Error())
		cobra.CheckErr(err)
	}

	err = exec.Command("rm", filepath.Join(env.BinaryDir(), "yoake")).Start()
	if err != nil {
		env.WriteLog(err.Error())
		cobra.CheckErr(err)
	}
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
