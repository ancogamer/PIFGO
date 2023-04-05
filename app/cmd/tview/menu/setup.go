package menu

import (
	"encoding/json"
	"fmt"

	"github.com/ancogamer/app/funcs"
	"github.com/rivo/tview"
)

type Setup struct {
	App *tview.Application
}

func (s *Setup) FileSave() *tview.List {
	list := tview.NewList()
	list.AddItem("Save file(s) Local", "", 'l', func() {
		s.App.SetRoot(s.setupDataLocal(), true)
	})
	// TODO
	list.AddItem("Save file(s) in Google drive", "", 'g', func() {
		fmt.Fprintln(tview.NewTextView(), "NOT SUPPORTED YET")
	})
	list.AddItem("Return", "Press r to return last menu", 'r', func() {
		s.App.SetRoot(Main(s.App), true)
	})
	return list
}

func (s *Setup) setupDataLocal() *tview.List {
	list := tview.NewList()
	list.AddItem("Use Current", "Start with current Data (uses System Data) and uses 1 year as default data window", 'c', func() {
		b, err := json.Marshal(funcs.SetupFileCurrentTime())
		if err != nil {
			fmt.Fprintln(tview.NewTextView(), "ERROR WHILE TRANSFORMING PROGRAM STRUCT", err)
		}
		err = funcs.WriteLocalFile("../../storage/data.json", b)
		if err != nil {
			fmt.Fprintln(tview.NewTextView(), "ERROR WHILE WRITING FILES", err)
		}
		fmt.Fprintln(tview.NewTextView(), "FILES SAVED")
	})

	// TODO
	list.AddItem("Manual Insert", "Manually insert start data", 'm', func() {
		fmt.Fprintln(tview.NewTextView(), "NOT SUPPORTED YET")
	})

	list.AddItem("Return", "Press r to return last menu", 'r', func() {
		s.App.SetRoot(s.FileSave(), true)
	})
	return list
}
