package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/mitsimi/yoake/cmd"
	"github.com/mitsimi/yoake/env"
)

func TestApp(t *testing.T) {
	// Run the app
	cmd.ExecuteArgs([]string{"init"})

	_os := runtime.GOOS
	// Check if STARTUP directory and file  created
	switch _os {
	case "windows":
		dir := env.Win_StartupDir()

		if _, err := os.Stat(env.Win_StartupDir()); os.IsNotExist(err) {
			t.Errorf("Windows startup directory not found, %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "yoake.bat")); os.IsNotExist(err) {
			t.Errorf("Windows script not found, %v", err)
		}
	case "linux":
		dir := env.Linux_StartupDir()

		if _, err := os.Stat(env.Linux_StartupDir()); os.IsNotExist(err) {
			t.Errorf("Linux startup directory not found, %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "yoake.desktop")); os.IsNotExist(err) {
			t.Errorf("Linux script not found, %v", err)
		}
	}

	// Check if CONFIG directory and file created
	if _, err := os.Stat(env.ConfigDir()); os.IsNotExist(err) {
		t.Errorf("Config directory not found, %v", err)
	}
	if _, err := os.Stat(env.ConfigFile()); os.IsNotExist(err) {
		t.Errorf("Config file not found, %v", err)
	}

}
