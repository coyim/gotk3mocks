package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockComboBox struct {
	MockBin
}

func (*MockComboBox) GetActiveIter() (gtki.TreeIter, error) {
	return nil, nil
}

func (*MockComboBox) GetActiveID() string {
	return ""
}

func (*MockComboBox) GetActive() int {
	return 0
}

func (*MockComboBox) SetActive(v1 int) {
}

func (*MockComboBox) SetModel(v1 gtki.TreeModel) {
}

func (*MockComboBox) AddAttribute(v1 gtki.CellRenderer, v2 string, v3 int) {
}

func (*MockComboBox) PackStart(v1 gtki.CellRenderer, v2 bool) {
}

func (*MockComboBox) SetIDColumn(v1 int) {
}

func (*MockComboBox) SetEntryTextColumn(v1 int) {
}

func (*MockComboBox) GetToggleButton() (gtki.Button, error) {
	return nil, nil
}
