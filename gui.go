package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func runGUI() error {
	if err := guiMain(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func guiMain() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-5, maxY/2-1, maxX/2+5, maxY/2+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, "Hello World!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
