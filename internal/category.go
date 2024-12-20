package internal

import (
	"fmt"
	"github.com/oliversabler/aid/style"
)

type Command struct {
	Keys        string
	Description string
}

type Category struct {
	ID       string
	Name     string
	Commands []Command
}

func GetCategoryIDs(categories []Category) []string {
	var categoryIDs []string
	for _, category := range categories {
		categoryIDs = append(categoryIDs, category.ID)
	}
	return categoryIDs
}

func GetCategoryCommands(categories []Category, categoryId string) Category {
	for _, tc := range categories {
		if tc.ID == categoryId {
			return tc
		}
	}

	return Category{}
}

func PrintCategoryIDs(categoryIds []string) {
	fmt.Println("")
	fmt.Printf("%s\n", style.Category.Render("Category IDs"))
	for _, id := range categoryIds {
		fmt.Printf("> %s\n", style.CategoryIDs.Render(id))
	}
	fmt.Println("")
}

func PrintCategories(categories []Category) {
	for _, category := range categories {
		PrintCategory(category)
	}
}

func PrintCategory(category Category) {
	fmt.Println("")
	fmt.Printf("%s\n", style.Category.Render(category.Name))
	for _, cmd := range category.Commands {
		printCommand(cmd.Keys, cmd.Description)
	}
	fmt.Println("")
}

func printCommand(command string, description string) {
	fmt.Printf("%s %s\n",
		style.Command.Render(command),
		style.Description.Render(description))
}
