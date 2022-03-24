package gtk

import (
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockTextTagTable struct {
	glib.MockObject
}

func (m *MockTextTagTable) Add(v1 gtki.TextTag) {
	m.Called(v1)
}
