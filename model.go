package main

import (
	list "charm.land/bubbles/v2/list"
)

type Model struct {
	layoutList     list.Model
	sizeList       list.Model
	activeLayout   string
	activeSize     int
	showLayoutList bool
	showSizeList   bool
	showInfoBar    bool
}

func getInitModel() Model {
	initModel := Model{
		layoutList:     list.New(keyboardLayoutItems, list.NewDefaultDelegate(), 0, 0),
		sizeList:       list.New(keyboardSizeItems, list.NewDefaultDelegate(), 0, 0),
		activeLayout:   "qwerty",
		activeSize:     75,
		showLayoutList: false,
		showSizeList:   false,
		showInfoBar:    true,
	}
	initModel.layoutList.Title = "Layouts"
	initModel.layoutList.KeyMap.Quit.SetKeys("q")
	initModel.sizeList.Title = "Sizes"
	initModel.sizeList.KeyMap.Quit.SetKeys("q")

	return initModel
}
