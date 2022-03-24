package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockMenuItem struct {
	MockBin
}

func (m *MockMenuItem) GetLabel() string {
	return m.Called().String(0)
}

func (m *MockMenuItem) SetLabel(v1 string) {
	m.Called(v1)
}

func (m *MockMenuItem) SetSubmenu(v1 gtki.Widget) {
	m.Called(v1)
}
