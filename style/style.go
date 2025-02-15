package style

import "github.com/charmbracelet/lipgloss"

var Category = lipgloss.NewStyle().
	Bold(true).
	Underline(true)

var CategoryIDs = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#33ffb5"))

var Command = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#f4ebba")).
	PaddingRight(2)

var Description = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#bad2f4"))
