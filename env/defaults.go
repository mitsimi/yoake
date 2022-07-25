package env

import (
	"fmt"
	"os"
	"runtime"

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
	home, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err)
	}

	switch runtime.GOOS {
	case "windows": // Windows
		return fmt.Sprintf("%s\\starigo", home)
	case "linux": // Linux
		return fmt.Sprintf("%s/starigo", home)
	}
	return ""
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
	return fmt.Sprintf(
		"Set WshShell = CreateObject(\"WScript.Shell\")\n"+
			"WshShell.Run chr(34) & \"%s\\starigo.exe\" & Chr(34), 0\n"+
			"Set WshShell = Nothing", BinaryDir(),
	)
}

func Linux_StartupDir() string {
	return fmt.Sprintf("%s\\.config\\autostart", UserHomeDir())
}

func Linux_Desktop() string {
	return fmt.Sprintf(
		"[Desktop Entry]\n"+
			"Type=Application\n"+
			"Name=StariGo\n"+
			"Exec=%s\\starigo\n"+
			"StartupNotify=false\n"+
			"Terminal=false", BinaryDir(),
	)
}
