package cmd

import (
	"github.com/oliversabler/aid/internal"
	"github.com/spf13/cobra"
)

var tmuxCategories = []internal.Category{
	{
		Name: "Windows",
		Commands: []internal.Command{
			internal.Command{Keys: "Ctrl + b c", Description: "Create new window"},
			internal.Command{Keys: "Ctrl + b ,", Description: "Rename current window"},
			internal.Command{Keys: "Ctrl + b &", Description: "Close current window"},
			internal.Command{Keys: "Ctrl + b n", Description: "Go to next window"},
			internal.Command{Keys: "Ctrl + b p", Description: "Go to previous window"},
			internal.Command{Keys: "Ctrl + b w", Description: "List all windows"},
			internal.Command{Keys: "Ctrl + b 0..9", Description: "Go to window by number"},
			internal.Command{Keys: "Ctrl + b l", Description: "Toggle last active window"},
		},
	},
	{
		Name: "Panes",
		Commands: []internal.Command{
			internal.Command{Keys: "Ctrl + b %", Description: "Create new pane with vertical line"},
			internal.Command{Keys: "Ctrl + b \"", Description: "Create new pane with horizontal line"},
			internal.Command{Keys: "Ctrl + b }", Description: "Move current pane right"},
			internal.Command{Keys: "Ctrl + b {", Description: "Move current pane left"},
			internal.Command{Keys: "Ctrl + b o", Description: "Switch to next pane"},
			internal.Command{Keys: "Ctrl + b q", Description: "Show pane numbers"},
			internal.Command{Keys: "Ctrl + b q 0..9", Description: "Go to pane by number"},
			internal.Command{Keys: "Ctrl + b z", Description: "Toggle pane zoom"},
			internal.Command{Keys: "Ctrl + b x", Description: "Close current pane"},
		},
	},
}

var tmuxCmd = &cobra.Command{
	Use:   "tmux",
	Short: "List tmux cmds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PrintCategories(tmuxCategories)
	},
}

func init() {
	rootCmd.AddCommand(tmuxCmd)
}
