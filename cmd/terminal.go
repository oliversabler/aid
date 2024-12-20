package cmd

import (
	"github.com/oliversabler/aid/internal"
	"github.com/spf13/cobra"
)

var terminalCategories = []internal.Category{
	{
		ID:   "movement",
		Name: "Movement",
		Commands: []internal.Command{
			{Keys: "Ctrl + A", Description: "Move cursor to the start of the line"},
			{Keys: "Ctrl + E", Description: "Move cursor to the end of the line"},
		},
	},
	{
		ID:   "deletion",
		Name: "Deletion",
		Commands: []internal.Command{
			{Keys: "Ctrl + L", Description: "Clear screen"},
			{Keys: "Ctrl + U", Description: "Erase the whole line"},
			{Keys: "Ctrl + K", Description: "Erase from cursor to the end of the line"},
			{Keys: "Ctrl + W", Description: "Erase the word before the cursor position"},
		},
	},
	{
		ID:   "commands",
		Name: "Commands",
		Commands: []internal.Command{
			{Keys: "Ctrl + P", Description: "Go to previous command"},
			{Keys: "Ctrl + R", Description: "Search for a command in history"},
			{Keys: "!!      ", Description: "Repeat last command"},
		},
	},
}

var terminalCmd = &cobra.Command{
	Use:   "terminal",
	Short: "List terminal cmds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listCategoryIDs, _ := cmd.Flags().GetBool("list-category-ids")
		if listCategoryIDs {
			categoryIDs := internal.GetCategoryIDs(terminalCategories)
			internal.PrintCategoryIDs(categoryIDs)
		} else {
			category, _ := cmd.Flags().GetString("category")

			if category != "" {
				category := internal.GetCategoryCommands(terminalCategories, category)
				internal.PrintCategory(category)
			} else {
				internal.PrintCategories(terminalCategories)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(terminalCmd)

	var lc bool
	terminalCmd.PersistentFlags().BoolVarP(&lc, "list-category-ids", "l", false, "List all category ids")
	terminalCmd.PersistentFlags().StringP("category", "c", "", "List commands in a specific category")
}
