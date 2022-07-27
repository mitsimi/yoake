package cmd

import (
	"fmt"
	"strconv"

	"github.com/mitsimi/starigo/env"
	"github.com/spf13/cobra"
)

var delayCmd = &cobra.Command{
	Use:   "delay [time]",
	Short: "Show or set the delay before application start up",
	Long:  `Show the current delay before application start up or set the delay to a new value in seconds.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			newDelay, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("\033[31;1mInvalid delay value\033[0m")
				return
			}
			config.Conf.Delay = newDelay

			if err := config.WriteConfig(); err != nil {
				env.WriteLog(err.Error())
				cobra.CheckErr(err)
			}
		}

		fmt.Printf("Delay is set to \033[93m%v\033[0m seconds.", config.Conf.Delay)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(delayCmd)
}
