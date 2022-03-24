package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockCSSClassCellRenderer struct {
	MockCellRenderer
}

func (m *MockCSSClassCellRenderer) SetReal(v1 gtki.CellRenderer) {
	m.Called(v1)
}
