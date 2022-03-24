package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockButton struct {
	MockBin
}

func (m *MockButton) SetImage(v1 gtki.Widget) {
	m.Called(v1)
}

func (m *MockButton) GetLabel() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockButton) SetLabel(v1 string) {
	m.Called(v1)
}
