package menu

import (
	"strconv"
	"strings"
	"time"

	"github.com/ancogamer/app/funcs"
	"github.com/ancogamer/app/structs/tax"

	"github.com/ancogamer/app/structs"

	"github.com/rivo/tview"
)

func (t *TUI) StandardOps() {
	t.ops.Clear()
	t.isOpsStantard = true
	t.ops.AddItem("New Transaction", "Inserts a new transaction", 'n', func() {
		t.NewTransaction()
	})

}

func (t *TUI) NewTransaction() {
	t.ops.Clear()
	t.ops.AddItem("Income", "", 'i', func() {
		form := tview.NewForm()
		form.AddInputField("DATA ON YYYY/MM/DD", funcs.CurrentData().Format(structs.TimeFormatYYYYMMDD), 20, nil, nil) //TODO get current item from the in memory slice
		form.AddDropDown("Type", []string{"Investment FIX RETURN.", "SALARY"}, 0, nil)                                 // TODO get structs from operation.Income
		form.AddInputField("VALUE", "", 20, nil, nil)                                                                  // call function of value
		form.AddCheckbox("PAID", false, nil)                                                                           // is this paid ? (since operations can be on credit to be paid later)
		form.AddButton("Save", func() {
			t.viewerScreen = t.viewer
			t.ReloadNavigate()
			data := form.GetFormItem(0).(*tview.InputField)
			tData, err := time.Parse(structs.TimeFormatYYYYMMDD, data.GetText())
			if err != nil {
				t.viewer.SetText(strings.Join([]string{"ERROR WHILE TRANSFORMING DATA :", err.Error()}, ""))
				return
			}

			p := funcs.GetPositionOperationYearItem(&t.Data.TimeWindow, strconv.Itoa(tData.Year()), int(tData.Month()), int(tData.Day()))
			if p == nil {
				t.viewer.SetText(strings.Join([]string{"ERROR POSITION OF DATA IN RECORDS NOT FOUND\nCHECK FILES WITH PARAMS :",
					strconv.Itoa(tData.Year()),
					strconv.Itoa(int(tData.Month())),
					strconv.Itoa(int(tData.Day())),
				}, ""))
				return
			}

			inp := operation.Income{
				InvestmentsFix: &operation.FixReturn{
					Name: "TESTE SAVE",
					TaxYear: tax.Tax{
						Porcentage: 0.0,
					},
					TaxMonth: tax.Tax{
						Porcentage: 0.0,
					},
				},
				Value: 0.0,
			}
			funcs.AddOrUpdateIncome(
				&t.Data.TimeWindow[p[0]],
				&t.Data.TimeWindow[p[0]].Months[p[1]],
				&t.Data.TimeWindow[p[0]].Months[p[1]].Days[p[2]],
				inp,
			)
			//var err error
			//s.TUI.Data.BeginValue, err = strconv.ParseFloat(formIt.GetText(), 64)
			//if err != nil {
			//	s.TUI.viewerScreen = s.TUI.viewer
			//	s.TUI.viewer.SetText(strings.Join([]string{"ERROR WHILE TRANSFORMING PROGRAM STRUCT :", err.Error()}, ""))
			//	return
			//}
		})
		form.AddButton("Cancel", func() {
			t.viewerScreen = t.viewer
			t.App.SetFocus(t.viewerScreen)
			t.ReloadNavigate()
			t.NewTransaction()
		})
		form.SetBorder(true).SetTitle("Enter data").SetTitleAlign(tview.AlignLeft)

		t.viewerScreen = form
		t.App.SetFocus(t.viewerScreen)
		t.ReloadNavigate()
	})
	// TODO
	t.ops.AddItem("Outcome", "", 'o', func() {
		app := tview.NewApplication()
		form := tview.NewForm().
			AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
			AddInputField("First name", "", 20, nil, nil).
			AddInputField("Last name", "", 20, nil, nil).
			AddTextArea("Address", "", 40, 0, 0, nil).
			AddTextView("Notes", "This is just a demo.\nYou can enter whatever you wish.", 40, 2, true, false).
			AddCheckbox("Age 18+", false, nil).
			AddPasswordField("Password", "", 10, '*', nil).
			AddButton("Save", nil).
			AddButton("Quit", func() {
				app.Stop()
			})
		form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	})
}
