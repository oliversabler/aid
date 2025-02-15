package cmd

import (
	"github.com/oliversabler/aid/internal"
	"github.com/spf13/cobra"
)

var miscCategories = []internal.Category{
	{
		ID:   "combos",
		Name: "Combos",
		Commands: []internal.Command{
			{
				Keys:        "find . -name '*.<file extension>' | xargs -I {} cat {} | wc -l",
				Description: "Count lines of all files of type",
			},
		},
	},
}

var miscCmd = &cobra.Command{
	Use:   "misc",
	Short: "List misc cmds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listCategoryIDs, _ := cmd.Flags().GetBool("list-category-ids")
		if listCategoryIDs {
			categoryIDs := internal.GetCategoryIDs(miscCategories)
			internal.PrintCategoryIDs(categoryIDs)
		} else {
			category, _ := cmd.Flags().GetString("category")

			if category != "" {
				category := internal.GetCategoryCommands(miscCategories, category)
				internal.PrintCategory(category)
			} else {
				internal.PrintCategories(miscCategories)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(miscCmd)

	var lc bool
	miscCmd.PersistentFlags().BoolVarP(&lc, "list-category-ids", "l", false, "List all category ids")
	miscCmd.PersistentFlags().StringP("category", "c", "", "List commands in a specific category")
}
