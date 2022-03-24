package gtk

import (
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockSizeGroup struct {
	glib.MockObject
}

func (m *MockSizeGroup) SetMode(v1 gtki.SizeGroupMode) {
	m.Called(v1)
}
