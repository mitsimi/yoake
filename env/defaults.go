package env

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UserHomeDir() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return home
}

func ConfigDir() string {
	home, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%s\\sg", home)
}

func ConfigFile() string {
	return fmt.Sprintf("%s\\config.json", ConfigDir())
}

func BinaryDir() string {
	return fmt.Sprintf("%s\\go\\bin", UserHomeDir())
}

func LogFile() string {
	return fmt.Sprintf("%s\\startup.log", ConfigDir())
}

func Win_StartupDir() string {
	return fmt.Sprintf("%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup", UserHomeDir())
}

func Win_Script() string {
	return fmt.Sprintf("Set WshShell = CreateObject(\"WScript.Shell\")\nWshShell.Run chr(34) & \"%s\\starigo.exe\" & Chr(34), 0\nSet WshShell = Nothing", BinaryDir())
}

func Linux_StartupDir() string {
	return fmt.Sprintf("%s\\.config\\autostart", UserHomeDir())
}

func Linux_Desktop() string {
	return fmt.Sprintf("[Desktop Entry]\nType=Application\nName=StariGo\nExec=%s\\starigo\nStartupNotify=false\nTerminal=false", BinaryDir())
}
