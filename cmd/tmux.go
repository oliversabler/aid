package cmd

import (
	"github.com/oliversabler/aid/internal"
	"github.com/spf13/cobra"
)

var tmuxCategories = []internal.Category{
	{
		ID:   "windows",
		Name: "Windows",
		Commands: []internal.Command{
			{Keys: "Ctrl + b c", Description: "Create new window"},
			{Keys: "Ctrl + b ,", Description: "Rename current window"},
			{Keys: "Ctrl + b &", Description: "Close current window"},
			{Keys: "Ctrl + b n", Description: "Go to next window"},
			{Keys: "Ctrl + b p", Description: "Go to previous window"},
			{Keys: "Ctrl + b w", Description: "List all windows"},
			{Keys: "Ctrl + b 0..9", Description: "Go to window by number"},
			{Keys: "Ctrl + b l", Description: "Toggle last active window"},
		},
	},
	{
		ID:   "panes",
		Name: "Panes",
		Commands: []internal.Command{
			{Keys: "Ctrl + b %", Description: "Create new pane with vertical line"},
			{Keys: "Ctrl + b \"", Description: "Create new pane with horizontal line"},
			{Keys: "Ctrl + b }", Description: "Move current pane right"},
			{Keys: "Ctrl + b {", Description: "Move current pane left"},
			{Keys: "Ctrl + b o", Description: "Switch to next pane"},
			{Keys: "Ctrl + b q", Description: "Show pane numbers"},
			{Keys: "Ctrl + b q 0..9", Description: "Go to pane by number"},
			{Keys: "Ctrl + b z", Description: "Toggle pane zoom"},
			{Keys: "Ctrl + b x", Description: "Close current pane"},
		},
	},
}

var tmuxCmd = &cobra.Command{
	Use:   "tmux",
	Short: "List tmux cmds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listCategoryIDs, _ := cmd.Flags().GetBool("list-category-ids")
		if listCategoryIDs {
			categoryIDs := internal.GetCategoryIDs(tmuxCategories)
			internal.PrintCategoryIDs(categoryIDs)
		} else {
			category, _ := cmd.Flags().GetString("category")

			if category != "" {
				category := internal.GetCategoryCommands(tmuxCategories, category)
				internal.PrintCategory(category)
			} else {
				internal.PrintCategories(tmuxCategories)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(tmuxCmd)

	var lc bool
	tmuxCmd.PersistentFlags().BoolVarP(&lc, "list-category-ids", "l", false, "List all category ids")
	tmuxCmd.PersistentFlags().StringP("category", "c", "", "List commands in a specific category")
}
