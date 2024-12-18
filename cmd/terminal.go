package cmd

import (
	"github.com/oliversabler/aid/internal"
	"github.com/spf13/cobra"
)

var terminalCategories = []internal.Category{
	{
		Name: "Movement",
		Commands: []internal.Command{
			internal.Command{Keys: "Ctrl + A", Description: "Move cursor to the start of the line"},
			internal.Command{Keys: "Ctrl + E", Description: "Move cursor to the end of the line"},
		},
	},
	{
		Name: "Deletion",
		Commands: []internal.Command{
			internal.Command{Keys: "Ctrl + L", Description: "Clear screen"},
			internal.Command{Keys: "Ctrl + U", Description: "Erase the whole line"},
			internal.Command{Keys: "Ctrl + K", Description: "Erase from cursor to the end of the line"},
			internal.Command{Keys: "Ctrl + W", Description: "Erase the word before the cursor position"},
		},
	},
	{
		Name: "Commands",
		Commands: []internal.Command{
			internal.Command{Keys: "Ctrl + P", Description: "Go to previous command"},
			internal.Command{Keys: "Ctrl + R", Description: "Search for a command in history"},
			internal.Command{Keys: "!!      ", Description: "Repeat last command"},
		},
	},
}

var terminalCmd = &cobra.Command{
	Use:   "terminal",
	Short: "list terminal cmds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PrintCategories(terminalCategories)
	},
}

func init() {
	rootCmd.AddCommand(terminalCmd)
}
