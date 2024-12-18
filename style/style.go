package style

import "github.com/charmbracelet/lipgloss"

var Category = lipgloss.NewStyle().
	Bold(true).
	Underline(true)

var Command = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#f4ebba")).
	Width(16)

var Description = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#bad2f4"))
