package main

import (
	"fmt"
	"gocui-training/models"
	"strconv"
	"strings"

	"github.com/jroimartin/gocui"
)

var (
	bottom = false
	dView  *gocui.View
)

// relSize calculates the  sizes of the sites view width
// and the news view height in relation to the current terminal size
func relSize(g *gocui.Gui) (int, int) {
	tw, th := g.Size()

	return (tw * 3) / 10, (th * 70) / 100
}

func layout(g *gocui.Gui) error {
	// Get the current terminal size.
	tw, th := g.Size()

	// Get the relative size of the views
	rw, rh := relSize(g)

	// Title
	if v, err := g.SetView("new-user", 0, 0, rw, th-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[1]-New User"
	}

	// Name input
	if v, err := g.SetView("name", 1, 1, rw-1, 3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Highlight = true
		v.Title = "Name"

		if _, err = setCurrentViewOnTop(g, "name"); err != nil {
			return err
		}
	}

	// Age input
	if v, err := g.SetView("age", 1, 4, rw-1, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Highlight = true
		v.Title = "Age"
	}

	// Email input
	if v, err := g.SetView("email", 1, 7, rw-1, 9); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Highlight = true
		v.Title = "Email"
	}

	// User list view
	if v, err := g.SetView("user-list", rw+1, 0, tw-1, rh-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[2]-User list"
	}
	// Output view
	if v, err := g.SetView("debug", rw+1, rh+1, tw-1, th-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Debug Output"
		fmt.Fprintln(v, "Debug information will appear here...")
	}

	return nil
}

func submit(g *gocui.Gui, v *gocui.View) error {
	nameView, _ := g.View("name")
	ageView, _ := g.View("age")
	emailView, _ := g.View("email")

	name := strings.TrimSpace(nameView.Buffer())
	ageStr := strings.TrimSpace(ageView.Buffer())
	email := strings.TrimSpace(emailView.Buffer())

	if name == "" || ageStr == "" || email == "" {
		showMessage(g, "Error: All fields are required")
		return nil
	}

	age, err := strconv.Atoi(strings.TrimSpace(ageStr))
	if err != nil {
		showMessage(g, "Age must be a number")
		return nil
	}

	user := models.AddUser(name, age, email)
	showMessage(g, fmt.Sprintf("Added: %s (%d)", user.Name, user.Id))

	nameView.Clear()
	ageView.Clear()
	emailView.Clear()
	SwitchView(g, nameView)

	return nil
}

func SwitchView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(views)
	name := views[nextIndex]

	// reset cursor if view is empty
	if _, err := setCurrentViewOnTop(g, name); err != nil {
		showMessage(g, fmt.Sprintf("Error %s", err))
		return err
	}
	active = nextIndex

	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func showMessage(g *gocui.Gui, msg string) {
	g.Update(func(gui *gocui.Gui) error {
		v, err := g.View("debug")
		if err != nil {
			return err
		}
		fmt.Fprintln(v, msg)

		lines := len(v.BufferLines())
		if lines > 0 {
			_, viewHeight := v.Size()

			if viewHeight > 0 && lines > viewHeight {
				newOriginY := max(lines-viewHeight, 0)
				_ = v.SetOrigin(0, newOriginY)
			}
		}
		return nil
	})
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
