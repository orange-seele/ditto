package main

import (
	"fmt"
	"os"
	"strings"

	lipgloss "charm.land/lipgloss/v2"
)

func getView(m Model) string {
	termWidth, termHeight, err := getTerminalSize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	bg := lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("#0000FF")).
		TabWidth(0)

	layoutCommand := bg.Render("CTRL + L: Layout")
	sizeCommand := bg.Render("CTRL + S: Size")
	hideCommand := bg.Render("CTRL + H: Hide Bar")
	activeLayout := bg.Render(strings.ToUpper(m.keyboardLayout))
	activeSize := bg.Render(fmt.Sprintf("%d%%", m.keyboardSize))

	commands := lipgloss.JoinHorizontal(lipgloss.Bottom, layoutCommand, " ", sizeCommand, " ", hideCommand)
	activeConfigs := lipgloss.JoinHorizontal(lipgloss.Bottom, activeLayout, " ", activeSize)

	spacerWidth := termWidth - lipgloss.Width(activeConfigs) - lipgloss.Width(commands)
	spacer := strings.Repeat(" ", spacerWidth)

	bottomBar := lipgloss.JoinHorizontal(lipgloss.Bottom, activeConfigs, spacer, commands)
	s := lipgloss.Place(termWidth, termHeight, lipgloss.Left, lipgloss.Bottom, bottomBar)

	return s
}
