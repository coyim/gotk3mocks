package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockListBox struct {
	MockContainer
}

func (m *MockListBox) SelectRow(v1 gtki.ListBoxRow) {
	m.Called(v1)
}

func (m *MockListBox) GetRowAtIndex(v1 int) gtki.ListBoxRow {
	return ret[gtki.ListBoxRow](m.Called(v1), 0)
}

func (m *MockListBox) GetSelectedRow() gtki.ListBoxRow {
	return ret[gtki.ListBoxRow](m.Called(), 0)
}
