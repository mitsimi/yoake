package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

const (
	enabledStatus  = "Enabled"
	disabledStatus = "Disabled"
)

var status map[bool]string = map[bool]string{
	true:  enabledStatus,
	false: disabledStatus,
}

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
				t.AppendRow(table.Row{upperName(name), status[app.Enabled], app.Path})
			}
			t.SetStyle(table.StyleLight)
			t.SetAutoIndex(true)
			t.SortBy([]table.SortBy{
				{Name: "Status", Mode: table.Dsc},
				{Name: "Name", Mode: table.Asc},
			})
			t.Render()
			return
		}

		app, ok := config.Apps[args[0]]
		if !ok {
			fmt.Println("\033[31;1;3mApp not found!\033[0m")
			return
		}

		t.AppendRow(table.Row{upperName(args[0]), status[app.Enabled], app.Path})

		t.SetStyle(table.StyleLight)
		t.Render()
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func upperName(name string) string {
	var upperName strings.Builder
	for _, str := range strings.Split(name, " ") {
		upperName.WriteString(strings.ToUpper(str[0:1]) + str[1:])
		upperName.WriteString(" ")
	}
	return upperName.String()
}
