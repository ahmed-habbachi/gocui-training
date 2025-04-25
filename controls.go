package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func Keybindings(g *gocui.Gui) error {
	for i := range views {
		viewName := views[i] // Capture the current view name for the closure

		// Define the keybinding
		if err := g.SetKeybinding(viewName, gocui.KeyTab, gocui.ModNone, SwitchView); err != nil {
			return err
		}
	}

	// Submit on Enter from email view
	if err := g.SetKeybinding("email", gocui.KeyEnter, gocui.ModNone, submit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	return nil
}
