package env

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	ConfigName = "config"
	ConfigType = "json"
)

func UserHomeDir() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return home
}

func ConfigDir() string {
	dir, err := os.UserConfigDir()
	cobra.CheckErr(err)

	return filepath.Join(dir, "yoake")
}

func ConfigFile() string {
	return filepath.Join(ConfigDir(), ConfigName+"."+ConfigType)
}

func BinaryDir() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return filepath.Join(gopath, "bin")
}

func LogFile() string {
	return filepath.Join(ConfigDir(), "startup.log")
}

func Win_StartupDir() string {
	return filepath.Join(UserHomeDir(), "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
}

func Win_Script() string {
	path := filepath.Join(BinaryDir(), "yoake.exe")
	return fmt.Sprintf("@echo off\nstart %s", path)
}

func Linux_StartupDir() string {
	return filepath.Join(UserHomeDir(), ".config", "autostart")
}

func Linux_Desktop() string {
	path := filepath.Join(BinaryDir(), "yoake")
	return fmt.Sprintf(
		"[Desktop Entry]\n"+
			"Type=Application\n"+
			"Name=Yoake\n"+
			"Exec=%s\n"+
			"StartupNotify=false\n"+
			"Terminal=false",
		path,
	)
}
