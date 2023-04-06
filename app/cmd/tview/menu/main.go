package menu

import (
	"strings"

	"github.com/gdamore/tcell/v2"

	"github.com/ancogamer/app/funcs"
	"github.com/rivo/tview"
)

type TUI struct {
	App    *tview.Application
	viewer *tview.TextView
	source *tview.List
	ops    *tview.List
}

func (t *TUI) Main() tview.Primitive {
	//welcomeTab := tview.NewTextView().SetText(strings.Join([]string{"WELCOME TO PIFGO\n", "SYSTEM UP FROM :", structs.StartData.String()}, "")).SetBorder(true).SetTitle("Welcome")
	t.source = tview.NewList()
	t.viewer = tview.NewTextView()
	t.viewer.SetBorder(true).SetTitle("Viewer")

	// init case file not found
	_, err := funcs.ReadLocalFile("../../storage/data.json")
	if err != nil {
		if err.Error() == `open ../../storage/data.json: no such file or directory` {
			t.viewer.SetText("NO DATA FILE FOUND, PLEASE START A NEW ONE")
			s := Setup{App: t.App, TUI: t}
			s.FileSave()
		} else {
			t.viewer.SetText(strings.Join([]string{"ERROR WHILE READING FILES FOR BOOT :", err.Error()}, ""))
			return nil
		}
	}

	t.source.SetBorder(true).SetTitle("Sources [CTRL - I]")
	t.ops = tview.NewList()
	t.ops.SetBorder(true).SetTitle("Operations [CTRL - O]")
	t.ops.AddItem("New Transaction", "Inserts a new transaction", 'n', func() {
		t.viewer.SetText("TODO OPERATION NOT READY")
	})

	navigate := tview.NewGrid().SetRows(0, 0, 0).
		//func(p tview.Primitive, row int, column int, rowSpan int, colSpan int, minGridHeight int, minGridWidth int, focus bool)
		//AddItem(welcomeTab, 0, 0, 1, 1, 0, 0, false).
		AddItem(t.source, 0, 0, 1, 1, 0, 0, true).
		AddItem(t.ops, 1, 0, 2, 1, 0, 0, false)

	t.keyboard()

	return tview.NewGrid().
		SetRows(0, 2).
		SetColumns(40, 0).
		SetBorders(false).
		AddItem(navigate, 0, 0, 1, 1, 0, 0, true).
		AddItem(t.viewer, 0, 1, 1, 1, 0, 0, false)

}

// https://github.com/KenanBek/dbui/blob/6cad7ca8699b725df4bac2885e9096af31c838be/internal/tui/keyboard.go#L8
func (t *TUI) keyboard() {
	t.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlI:
			t.App.SetFocus(t.source)
		case tcell.KeyCtrlO:
			t.App.SetFocus(t.ops)
		}
		return event
	})
}
