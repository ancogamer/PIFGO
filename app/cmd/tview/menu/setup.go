package menu

import (
	"encoding/json"
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
}

func (s *Setup) setupDataLocal() {
	s.source.Clear()
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
