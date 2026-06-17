package main

import (
	list "charm.land/bubbles/v2/list"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

var keyboardLayoutItems = []list.Item{
	item{title: "QWERTY", desc: "The default"},
	item{title: "DVORAK", desc: "For nerds"},
	item{title: "COLEMAK", desc: "For mega nerds"},
	item{title: "COLEMAK-DH", desc: "Modern variant with better thumb positioning"},
	item{title: "WORKMAN", desc: "Reduces lateral finger movement"},
	item{title: "AZERTY", desc: "Standard French layout"},
}

var keyboardSizeItems = []list.Item{
	item{title: "60%", desc: "No backtick :("},
	item{title: "65%", desc: "Really?"},
	item{title: "75%", desc: "The classic"},
	item{title: "80%", desc: "Oooh"},
	item{title: "100%", desc: "Damn"},
}
