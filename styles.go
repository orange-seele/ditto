package main

import lipgloss "charm.land/lipgloss/v2"

var (
	listFrameStyle = lipgloss.NewStyle().Margin(1, 2)
	listStyle      = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2)
	bgStyle = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("#0000FF")).
		TabWidth(0)
)
