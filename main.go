package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

var (
	views  = []string{"name", "age", "email"}
	active = 0
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true

	g.SetManagerFunc(layout)

	// Force a layout update before setting keybindings
	g.Update(func(*gocui.Gui) error { return nil })

	// Apply keybindings to program.
	if err = Keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
