package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/viper"
)

func WriteLog(str string) {
	// Write log message to file
	time := fmt.Sprintf("[%s] ", time.Now().Format("02/01/2006 15:04:05"))
	message := fmt.Sprintf("%s%s\n", time, str)

	logFile := LogFile()

	// Check if log file exists
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		// Create log file
		err := ioutil.WriteFile(logFile, []byte(message), 0644)
		if err != nil {
			panic(err)
		}
		return
	}

	// Append log message to file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteString(message); err != nil {
		panic(err)
	}
}

func LoadConfig() (config Configuration, err error) {
	// Check if config directory exists
	if _, err := os.Stat(ConfigDir()); os.IsNotExist(err) {
		// Create config directory
		err := os.Mkdir(ConfigDir(), 0755)
		if err != nil {
			WriteLog(err.Error())
		}
	}

	// Configure viper
	viper.AddConfigPath(ConfigDir())
	viper.SetConfigType(ConfigType)
	viper.SetConfigName(ConfigName)

	viper.SetDefault("config", Config{
		Enabled: true,
		Delay:   10,
	})
	viper.SetDefault("apps", map[string]App{})

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			WriteLog(err.Error())
			WriteLog(fmt.Sprintf("Creating New Config File \"%s.%s\"", ConfigName, ConfigType))

			// Write config file
			err := ioutil.WriteFile(ConfigFile(), nil, 0644)
			if err != nil {
				WriteLog(err.Error())
				return config, err
			}
			viper.WriteConfig()
			if err != nil {
				WriteLog(err.Error())
				return config, err
			}
			WriteLog("Config File Created")

		} else {
			// Config file was found but another error was produced
			WriteLog(err.Error())
			return config, err
		}
	}
	// Config file ready
	err = viper.Unmarshal(&config)
	if err != nil {
		WriteLog("Unable To Decode Into Struct, %v" + err.Error())
		return config, err
	}
	return
}

func (conf *Configuration) WriteConfig() (err error) {
	viper.GetViper().Set("config", conf.Conf)
	viper.GetViper().Set("apps", conf.Apps)

	err = viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
