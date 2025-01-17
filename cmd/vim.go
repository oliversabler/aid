package cmd

import (
	"github.com/oliversabler/aid/internal"
	"github.com/spf13/cobra"
)

var vimCategories = []internal.Category{
	{
		ID:   "movement",
		Name: "Movement",
		Commands: []internal.Command{
			{Keys: "}", Description: "Jump to next paragraph, function or block"},
			{Keys: "{", Description: "Jump to previous paragraph, function or block"},
			{Keys: "%", Description: "Move cursor to matching character - '()', '{}', '[]'"},
		},
	},
	{
		ID:   "marking",
		Name: "Marking text (visual mode)",
		Commands: []internal.Command{
			{Keys: "aw", Description: "Mark a word"},
			{Keys: "ab", Description: "Mark a block with ()"},
			{Keys: "aB", Description: "Mark a block with {}"},
			{Keys: "at", Description: "Mark a block with <>"},
			{Keys: "ib", Description: "Mark text inside a block with ()"},
			{Keys: "iB", Description: "Mark text inside a block with {}"},
			{Keys: "it", Description: "Mark text inside a block with <>"},
		},
	},
	{
		ID:   "custom",
		Name: "Custom (visual mode)",
		Commands: []internal.Command{
			{Keys: "ctrl + k", Description: "Move current line up one line"},
			{Keys: "ctrl + j", Description: "Move current line down one line"},
		},
	},
	{
		ID:   "oil",
		Name: "Plugin (oil)",
		Commands: []internal.Command{
			{Keys: "-", Description: "Open parent directory"},
			{Keys: "ctrl + p", Description: "Open file preview"},
			{Keys: "gs", Description: "Change sorting"},
			{Keys: "gx", Description: "Open external"},
		},
	},
}

var vimCmd = &cobra.Command{
	Use:   "vim",
	Short: "List vim cmds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listCategoryIDs, _ := cmd.Flags().GetBool("list-category-ids")
		if listCategoryIDs {
			categoryIDs := internal.GetCategoryIDs(vimCategories)
			internal.PrintCategoryIDs(categoryIDs)
		} else {
			category, _ := cmd.Flags().GetString("category")

			if category != "" {
				category := internal.GetCategoryCommands(vimCategories, category)
				internal.PrintCategory(category)
			} else {
				internal.PrintCategories(vimCategories)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(vimCmd)

	var lc bool
	vimCmd.PersistentFlags().BoolVarP(&lc, "list-category-ids", "l", false, "List all category ids")
	vimCmd.PersistentFlags().StringP("category", "c", "", "List commands in a specific category")
}
