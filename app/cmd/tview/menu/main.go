package menu

import (
	"strings"

	"github.com/ancogamer/app/funcs"
	"github.com/rivo/tview"
)

func Main(app *tview.Application) tview.Primitive {
	source := tview.NewList()
	var view tview.Primitive
	tv := tview.NewTextView()
	tv.SetBorder(true).SetTitle("Viewer")

	// TODO WELCOME VIEW
	//fmt.Fprintln(tview.NewTextView(), "WELCOME TO PIFGO")
	//fmt.Fprintln(tview.NewTextView(), "SYSTEM UP FROM :", structs.StartData.String())

	// init case file not found
	_, err := funcs.ReadLocalFile("../../storage/data.json")
	if err != nil {
		view = tv
		if err.Error() == `open ../../storage/data.json: no such file or directory` {
			tv.SetText("NO DATA FILE FOUND, PLEASE START A NEW ONE")
			s := Setup{App: app}
			source = s.FileSave()
		} else {
			tv.SetText(strings.Join([]string{"ERROR WHILE READING FILES FOR BOOT :", err.Error()}, ""))
			return nil
		}
	}
	source.SetBorder(true).SetTitle("Sources")
	operations := tview.NewList()
	operations.SetBorder(true).SetTitle("Operations")
	operations.AddItem("New Transaction", "Inserts a new transaction", 'n', func() {
		tv.SetText("TODO OPERATION NOT READY")
	})

	navigate := tview.NewGrid().SetRows(0, 0, 0).
		//func(p tview.Primitive, row int, column int, rowSpan int, colSpan int, minGridHeight int, minGridWidth int, focus bool)
		AddItem(source, 0, 0, 1, 1, 0, 0, true).
		AddItem(operations, 1, 0, 2, 1, 0, 0, false)

	return tview.NewGrid().
		SetRows(0, 2).
		SetColumns(40, 0).
		SetBorders(false).
		AddItem(navigate, 0, 0, 1, 1, 0, 0, true).
		AddItem(view, 0, 1, 1, 1, 0, 0, false)
}
