package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "starigo",
	Short:   "Little cli program for easy and universal application start on startup",
	Long:    `A simple yet intuitive cli tool for managing windows and linux start up applications.`,
	Version: "0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	os := runtime.GOOS

	switch os {
	case "windows": // Windows
		break
	case "linux": // Linux
		break
	default: // Other
		fmt.Println("Unsupported OS")
	}
}
