package menu

import (
	"github.com/ancogamer/app/funcs"

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

		form := tview.NewForm().
			AddInputField("DATA ON YYYY/MM/DD", funcs.CurrentData().Format(structs.TimeFormatYYYYMMDD), 20, nil, nil).
			AddDropDown("Type", []string{"Investment FIX RETURN.", "SALARY"}, 0, nil).
			AddInputField("VALUE", "", 20, nil, nil).
			AddCheckbox("PAID", false, nil).
			AddButton("Save", nil).
			AddButton("Cancel", func() {
				t.NewTransaction()
			})
		form.SetBorder(true).SetTitle("Enter data").SetTitleAlign(tview.AlignLeft)

		t.viewerScreen = form
		t.App.SetFocus(t.viewerScreen)

		t.ReloadNavigate()
	})
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
