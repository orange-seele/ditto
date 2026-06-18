package main

import (
	"image/color"
	"os"

	lipgloss "charm.land/lipgloss/v2"
)

var listFrameStyle = lipgloss.NewStyle().Margin(1, 2)

var (
	listStyle    lipgloss.Style
	fingerStyle  map[Finger]lipgloss.Style
	fingerActive map[Finger]lipgloss.Style
	infoBarStyle lipgloss.Style
)

var darkColors = map[Finger]color.Color{
	FingerPinky:  lipgloss.BrightMagenta,
	FingerRing:   lipgloss.BrightRed,
	FingerMiddle: lipgloss.BrightYellow,
	FingerIndex:  lipgloss.BrightBlue,
	FingerThumb:  lipgloss.BrightCyan,
}

var lightColors = map[Finger]color.Color{
	FingerPinky:  lipgloss.Magenta,
	FingerRing:   lipgloss.Red,
	FingerMiddle: lipgloss.Yellow,
	FingerIndex:  lipgloss.Blue,
	FingerThumb:  lipgloss.Cyan,
}

func init() {
	isDark := lipgloss.HasDarkBackground(os.Stdin, os.Stdout)

	colors := darkColors
	borderColor := lipgloss.BrightBlack

	if !isDark {
		colors = lightColors
		borderColor = lipgloss.Black
	}

	listStyle = lipgloss.NewStyle().
		BorderForeground(borderColor).
		Border(lipgloss.ThickBorder()).
		Padding(1, 2)

	fingerStyle = make(map[Finger]lipgloss.Style, len(colors))
	fingerActive = make(map[Finger]lipgloss.Style, len(colors))

	for finger, c := range colors {
		base := lipgloss.NewStyle().Foreground(c)

		if isDark {
			fingerStyle[finger] = base.Faint(true)
		} else {
			fingerStyle[finger] = base
		}

		fingerActive[finger] = base.Copy().Bold(true).Italic(true)
	}

	if isDark {
		infoBarStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	} else {
		infoBarStyle = lipgloss.NewStyle().Faint(true)
	}
}
