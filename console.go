package main

import (
	"fmt"
	"os"

	"github.com/awesome-gocui/gocui"
)

func guilayout(gui *gocui.Gui) error {
	return nil
}

func initKeybindings(gui *gocui.Gui) error {
	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(gui *gocui.Gui, guiv *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}

	if err := gui.SetKeybinding("", gocui.KeyEsc, gocui.ModNone,
		func(gui *gocui.Gui, guiv *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	return nil
}

func guiInit() {
	gui, err := gocui.NewGui(gocui.Output256, true)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(1)
	}
	defer gui.Close()
	gui.SetManagerFunc(guilayout)

	if err := initKeybindings(gui); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	x0 := 0
	y0 := 0
	x1 := 80 // (c8.ScreenSize[0] / 2) + 30
	y1 := 24 // c8.ScreenSize[1] / 2

	name := "Screen"
	guiv, err := gui.SetView(name, x0, y0, int(x1), int(y1), 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return
		}
	}
	guiv.Frame = true
	if _, err := gui.SetCurrentView(name); err != nil {
		return
	}

	go guiStart(gui, guiv)
	gui.Update(func(gui *gocui.Gui) error {
		return nil
	})
}

func guiStart(gui *gocui.Gui, guiv *gocui.View) {
	err := gui.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		fmt.Println(err)
		os.Exit(1)
	}
}
