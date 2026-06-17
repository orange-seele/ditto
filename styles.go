package main

import lipgloss "charm.land/lipgloss/v2"

var (
	listFrameStyle = lipgloss.NewStyle().Margin(1, 2)

	listStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.BrightBlack).
			Border(lipgloss.ThickBorder()).
			Padding(1, 2)

	fingerStyle = map[Finger]lipgloss.Style{
		FingerPinky:  lipgloss.NewStyle().Foreground(lipgloss.BrightMagenta),
		FingerRing:   lipgloss.NewStyle().Foreground(lipgloss.BrightRed),
		FingerMiddle: lipgloss.NewStyle().Foreground(lipgloss.BrightYellow),
		FingerIndex:  lipgloss.NewStyle().Foreground(lipgloss.BrightCyan),
		FingerThumb:  lipgloss.NewStyle().Foreground(lipgloss.BrightGreen),
	}
)
