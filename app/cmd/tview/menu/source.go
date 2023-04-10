package menu

func (t *TUI) StandardSource() {
	t.source.Clear()
	t.isSourceStantard = true
	t.source.AddItem("Load file(s) Local", "", 'l', func() {
		t.viewer.SetText("TODO OPERATION NOT READY")
	})
	// TODO
	t.source.AddItem("Load file(s) from Google drive", "", 'g', func() {
		t.viewer.SetText("TODO OPERATION NOT READY")
	})
}
