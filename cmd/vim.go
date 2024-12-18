package cmd

import (
	"github.com/oliversabler/aid/internal"
	"github.com/spf13/cobra"
)

var vimCategories = []internal.Category{
	{
		Name: "Movement",
		Commands: []internal.Command{
			internal.Command{Keys: "}", Description: "Jump to next paragraph, function or block"},
			internal.Command{Keys: "{", Description: "Jump to previous paragraph, function or block"},
			internal.Command{Keys: "%", Description: "Move cursor to matching character - '()', '{}', '[]'"},
		},
	},
	{
		Name: "Marking text (visual mode)",
		Commands: []internal.Command{
			internal.Command{Keys: "aw", Description: "Mark a word"},
			internal.Command{Keys: "ab", Description: "Mark a block with ()"},
			internal.Command{Keys: "aB", Description: "Mark a block with {}"},
			internal.Command{Keys: "at", Description: "Mark a block with <>"},
			internal.Command{Keys: "ib", Description: "Mark text inside a block with ()"},
			internal.Command{Keys: "iB", Description: "Mark text inside a block with {}"},
			internal.Command{Keys: "it", Description: "Mark text inside a block with <>"},
		},
	},
	{
		Name: "Custom (visual mode)",
		Commands: []internal.Command{
			internal.Command{Keys: "ctrl + k", Description: "Move current line up one line"},
			internal.Command{Keys: "ctrl + j", Description: "Move current line down one line"},
		},
	},
}

var vimCmd = &cobra.Command{
	Use:   "vim",
	Short: "List vim cmds",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PrintCategories(vimCategories)
	},
}

func init() {
	rootCmd.AddCommand(vimCmd)
}
