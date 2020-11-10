package gui

import (
	"github.com/TNK-Studio/lazykube/pkg/log"
	"github.com/jroimartin/gocui"
)

var (
	Quit = &Action{
		Name: "Quit",
		Key: gocui.KeyCtrlC,
		Handler: func(gui *Gui) func(*gocui.Gui, *gocui.View) error {
			return func(*gocui.Gui, *gocui.View) error {
				return gocui.ErrQuit
			}
		},
		Mod: gocui.ModNone,
	}

	ClickView = &Action{
		Name:    "clickView",
		Key:     gocui.MouseLeft,
		Handler: ViewClickHandler,
		Mod:     gocui.ModNone,
	}
)

type Action struct {
	Name string
	Key     interface{}
	Handler func(gui *Gui) func(*gocui.Gui, *gocui.View) error
	Mod     gocui.Modifier
}

type ActionHandler func(gui *Gui) func(*gocui.Gui, *gocui.View) error

func ViewClickHandler(gui *Gui) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		viewName := v.Name()
		log.Logger.Debugf("ViewClickHandler view '%s' on click.", viewName)

		currentView := gui.CurrentView()

		if currentView != nil && currentView.Name == viewName {
			return nil
		}

		view, err := gui.GetView(viewName)
		if err != nil {
			if err == gocui.ErrUnknownView {
				return nil
			}
			return err
		}

		if err := gui.FocusView(viewName, true); err != nil {
			return err
		}


		if view.OnClick != nil {
			return view.OnClick(gui, view)
		}

		return nil
	}
}