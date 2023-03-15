package cmd

import (
	"fmt"

	"github.com/priuatus/submind/app"
	"github.com/priuatus/submind/color"
	"github.com/priuatus/submind/filesystem"
	"github.com/priuatus/submind/icon"
	"github.com/priuatus/submind/style"
	"github.com/priuatus/submind/util"
	"github.com/priuatus/submind/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

type clearTarget struct {
	name  string
	clear func() error
}

// Specify what can be cleared
var clearTargets = []clearTarget{
	{"cache", func() error {
		return filesystem.Api().RemoveAll(where.Cache())
	}},
	{"logs", func() error {
		return filesystem.Api().RemoveAll(where.Logs())
	}},
}

func init() {
	rootCmd.AddCommand(clearCmd)
	for _, n := range clearTargets {
		clearCmd.Flags().BoolP(n.name, string(n.name[0]), false, "clear "+n.name)
	}
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears sidelined files produced by the " + app.Name,
	Run: func(cmd *cobra.Command, args []string) {
		successStyle := style.Fg(color.Green)
		var didSomething bool
		for _, n := range clearTargets {
			if lo.Must(cmd.Flags().GetBool(n.name)) {
				handleErr(n.clear())
				fmt.Printf("%s %s cleared\n", successStyle(icon.Check), util.Capitalize(n.name))
				didSomething = true
			}
		}

		if !didSomething {
			_ = cmd.Help()
		}
	},
}
