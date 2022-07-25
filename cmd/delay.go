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
				fmt.Println("Invalid delay value")
				return
			}
			config.Conf.Delay = newDelay

			if err := config.WriteConfig(); err != nil {
				env.WriteLog(err.Error())
				cobra.CheckErr(err)
			}
		}

		fmt.Println("Delay is set to", config.Conf.Delay)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(delayCmd)
}
