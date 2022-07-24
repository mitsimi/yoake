package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show [app]",
	Aliases: []string{"list", "ls"},
	Short:   "Lists all or a single app",
	Long:    `Displays information about either all added applications or a specified one.`,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Name", "Status", "Path"})

		if len(args) == 0 || args[0] == "all" {
			for name, app := range config.Apps {
				var status string
				if app.Enabled {
					status = "enabled"
				} else {
					status = "disabled"
				}
				t.AppendRow(table.Row{name, status, app.Path})
			}
			t.SetStyle(table.StyleLight)
			t.Render()
			return
		}

		app := config.Apps[args[0]]
		var status string
		if app.Enabled {
			status = "enabled"
		} else {
			status = "disabled"
		}
		t.AppendRow(table.Row{args[0], status, app.Path})

		t.SetStyle(table.StyleLight)
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
