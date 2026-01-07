package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var styleError = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#ff6b6b"))

var styleResponse = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a8e6cf"))

var styleInfo = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#888888")).
	Italic(true)

func PrintError(msg string) {
	fmt.Fprintln(os.Stderr, styleError.Render("Error: "+msg))
}

func PrintResponse(response string) {
	fmt.Println(styleResponse.Render(response))
}

func PrintInfo(msg string) {
	fmt.Println(styleInfo.Render(msg))
}
