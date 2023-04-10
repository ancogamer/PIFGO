package menu

import "github.com/rivo/tview"

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

		t.viewerScreen = tview.NewForm().
			AddButton("Use current day", nil).
			AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
			AddInputField("First name", "", 20, nil, nil).
			AddInputField("Last name", "", 20, nil, nil).
			AddTextArea("Address", "", 40, 0, 0, nil).
			AddTextView("Notes", "This is just a demo.\nYou can enter whatever you wish.", 40, 2, true, false).
			AddCheckbox("Age 18+", false, nil).
			AddPasswordField("Password", "", 10, '*', nil).
			AddButton("Save", nil).
			AddButton("Cancel", func() {
				t.NewTransaction()
			}).SetBorder(true).SetTitle("Enter data").SetTitleAlign(tview.AlignLeft)
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
