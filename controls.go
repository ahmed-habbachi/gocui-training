package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func keybindings(g *gocui.Gui) error {
	views := []string{"name", "age", "email"}

	// Tab through views
	for i := range views {
		next := views[(i+1)%len(views)]
		if err := g.SetKeybinding(views[i], gocui.KeyTab, gocui.ModNone, switchView(next)); err != nil {
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
