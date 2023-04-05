package main

import (
	"github.com/ancogamer/app/cmd/tview/menu"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	rootMenu := menu.Main(app)
	if err := app.SetRoot(rootMenu, true).SetFocus(rootMenu).Run(); err != nil {
		panic(err)
	}
}
