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

func layout(g *gocui.Gui) error {
	maxX, _ := g.Size()

	// Title
	if v, err := g.SetView("title", maxX/2-15, 1, maxX/2+15, 3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Add New User")
	}

	// Name input
	if v, err := g.SetView("name", maxX/2-20, 4, maxX/2+20, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Title = "Name"
	}

	// Age input
	if v, err := g.SetView("age", maxX/2-20, 7, maxX/2+20, 9); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Title = "Age"
	}

	// Email input
	if v, err := g.SetView("email", maxX/2-20, 10, maxX/2+20, 12); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Title = "Email"
	}

	// Output view
	if v, err := g.SetView("output", maxX/2-20, 13, maxX/2+20, 16); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Title = "Info"
	}

	// Set default view to Name
	if _, err := g.SetCurrentView("name"); err != nil {
		return err
	}

	return nil
}

func submit(g *gocui.Gui, v *gocui.View) error {
	input := strings.TrimSpace(v.Buffer())
	parts := strings.Split(input, ",")
	if len(parts) != 3 {
		showMessage(g, "Invalid input. Use: name,age,email")
		return nil
	}

	name := strings.TrimSpace(parts[0])
	age, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		showMessage(g, "Age must be a number")
		return nil
	}
	email := strings.TrimSpace(parts[2])

	user := models.AddUser(name, age, email)
	showMessage(g, fmt.Sprintf("Added: %s (%d)", user.Name, user.Id))

	v.Clear()
	v.SetCursor(0, 0)
	return nil
}

func switchView(next string) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		_, err := g.SetCurrentView(next)
		return err
	}
}

func showMessage(g *gocui.Gui, msg string) {
	g.Update(func(gui *gocui.Gui) error {
		v, err := g.View("output")
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, msg)
		return nil
	})
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
