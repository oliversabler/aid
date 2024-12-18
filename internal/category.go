package internal

import (
	"fmt"
	"github.com/oliversabler/aid/style"
)

type Command struct {
	Keys        string
	Description string
	Subcategory string
}

type Category struct {
	Name     string
	Commands []Command
}

func PrintCategories(categories []Category) {
	fmt.Println("")
	for i, category := range categories {
		if i != 0 {
			fmt.Println()
		}
		fmt.Printf("%s\n", style.Category.Render(category.Name))
		for _, cmd := range category.Commands {
			printCommand(cmd.Keys, cmd.Description)
		}
	}
	fmt.Println("")
}

func printCommand(command string, description string) {
	fmt.Printf("%s %s\n",
		style.Command.Render(command),
		style.Description.Render(description))

}
