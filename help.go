package main

import (
	key "charm.land/bubbles/v2/key"
)

type CommandMap struct {
	Layout  key.Binding
	Size    key.Binding
	HideKey key.Binding
	Help    key.Binding
	Quit    key.Binding
}

func (c CommandMap) ShortHelp() []key.Binding {
	return []key.Binding{c.Help, c.Quit}
}

func (c CommandMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{c.Layout},
		{c.Size},
		{c.HideKey},
		{c.Quit},
	}
}

var commands = CommandMap{
	Layout: key.NewBinding(
		key.WithKeys("ctrl+shift+l"),
		key.WithHelp("ctrl+shift+l", "layout"),
	),
	Size: key.NewBinding(
		key.WithKeys("ctrl+shift+s"),
		key.WithHelp("ctrl+shift+s", "size"),
	),
	HideKey: key.NewBinding(
		key.WithKeys("ctrl+shift+h"),
		key.WithHelp("ctrl+shift+h", "hide Bar"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
