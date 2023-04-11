package menu

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/ancogamer/app/funcs"
	"github.com/rivo/tview"
)

type Setup struct {
	*TUI
	App *tview.Application
}

func (s *Setup) FileSave() {
	s.source.Clear()
	form := tview.NewForm()

	form.AddInputField("Initial Value", "", 20, nil, nil)
	form.AddButton("Save", func() {
		s.TUI.viewer.Clear()
		s.TUI.viewerScreen = s.TUI.viewer
		s.App.SetFocus(s.source)
		s.ReloadNavigate()

		formIt := form.GetFormItem(0).(*tview.InputField)
		var err error
		s.TUI.Data.BeginValue, err = strconv.ParseFloat(formIt.GetText(), 64)
		if err != nil {
			s.TUI.viewerScreen = s.TUI.viewer
			s.TUI.viewer.SetText(strings.Join([]string{"ERROR WHILE TRANSFORMING PROGRAM STRUCT :", err.Error()}, ""))
			return
		}

		s.source.AddItem("Save file(s) Local", "", 'l', func() {
			s.setupDataLocal()
		})
		// TODO
		s.source.AddItem("Save file(s) in Google drive", "", 'g', func() {
			s.TUI.viewer.SetText("TODO OPERATION NOT READY")
		})
		s.source.AddItem("Return", "Press r to return last menu", 'r', func() {
			s.StandardSource()
		})
	})

	form.SetBorder(true).SetTitle("Enter data").SetTitleAlign(tview.AlignLeft)

	s.viewerScreen = form
	s.ReloadNavigate()
	s.App.SetFocus(s.viewerScreen)

}

func (s *Setup) setupDataLocal() {
	s.source.Clear()
	s.App.SetFocus(s.source)
	s.source.AddItem("Use Current", "Start with current Data (uses System Data) and uses 1 year as default data window", 'c', func() {
		s.TUI.Data.TimeWindow = funcs.SetupFileCurrentTime()
		b, err := json.Marshal(s.TUI.Data)
		if err != nil {
			s.TUI.viewer.SetText(strings.Join([]string{"ERROR WHILE TRANSFORMING PROGRAM STRUCT :", err.Error()}, ""))
			return
		}
		err = funcs.WriteLocalFile("../../storage/data.json", b)
		if err != nil {
			s.TUI.viewer.SetText(strings.Join([]string{"ERROR WHILE WRITING FILES :", err.Error()}, ""))
		}
		s.TUI.viewer.SetText("FILES SAVED AND LOADED INTO MEMORY")
		s.StandardSource()
		s.StandardOps()
	})

	// TODO
	s.source.AddItem("Manual Insert", "Manually insert start data", 'm', func() {
		s.TUI.viewer.SetText("TODO OPERATION NOT READY")
	})

	s.source.AddItem("Return", "Press r to return last menu", 'r', func() {
		s.FileSave()
	})

}
