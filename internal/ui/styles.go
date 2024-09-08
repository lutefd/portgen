package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	primaryColor    = lipgloss.Color("#4a4e69")
	secondaryColor  = lipgloss.Color("#9a8c98")
	accentColor     = lipgloss.Color("#c9ada7")
	backgroundColor = lipgloss.Color("#f2e9e4")
	successColor    = lipgloss.Color("#2ecc71")
	offColor        = lipgloss.Color("#e74c3c")

	TitleStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)
	ClipboardStatusStyle = lipgloss.NewStyle().
				Foreground(secondaryColor).
				Italic(true)

	InputStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).MarginTop(1)

	PortStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true).
			MarginTop(1)

	InfoStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Italic(true)

	ShortPortStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true).
			Padding(0, 1)

	HelpHeadingStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
				MarginTop(1).
				MarginBottom(1)

	HelpCommandStyle = lipgloss.NewStyle().
				Foreground(accentColor).
				Bold(true)

	HelpDescStyle = lipgloss.NewStyle().
			Foreground(secondaryColor)
	ErrorDescStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).MarginBottom(1)
	TempMessageStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true)
	ClipboardOnStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true)

	ClipboardOffStyle = lipgloss.NewStyle().
				Foreground(offColor).
				Bold(true)
)

func GetLongDescription() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		HelpHeadingStyle.Render("Portgen"),
		HelpDescStyle.Render("A CLI tool to generate a random unused port within a specified range."),
		HelpDescStyle.Render("The generated port can optionally be copied to the clipboard."),
	)
}

func GetUsageTemplate() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		HelpHeadingStyle.Render("Usage:"),
		HelpCommandStyle.Render("  portgen [flags]"),
		"",
		HelpHeadingStyle.Render("Flags:"),
		HelpCommandStyle.Render("  -m, --min int    ")+HelpDescStyle.Render("Minimum port number (inclusive) (default 10000)"),
		HelpCommandStyle.Render("  -M, --max int    ")+HelpDescStyle.Render("Maximum port number (inclusive) (default 65535)"),
		HelpCommandStyle.Render("  -c, --copy       ")+HelpDescStyle.Render("Copy the generated port to clipboard"),
		HelpCommandStyle.Render("  -s, --short      ")+HelpDescStyle.Render("Print only the generated port number"),
		HelpCommandStyle.Render("  -h, --help       ")+HelpDescStyle.Render("Display help for portgen"),
	)
}
