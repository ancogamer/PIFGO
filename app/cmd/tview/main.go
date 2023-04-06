package main

import (
	"github.com/ancogamer/app/cmd/tview/menu"
	"github.com/rivo/tview"
)

func main() {

	tui := menu.TUI{App: tview.NewApplication()}
	if err := tui.App.SetRoot(tui.Main(), true).Run(); err != nil {
		panic(err)
	}
}
