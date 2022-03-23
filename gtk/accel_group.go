package gtk

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockAccelGroup struct {
	glib.MockObject
}

func (m *MockAccelGroup) Connect2(v2 uint, v3 gdki.ModifierType, v4 gtki.AccelFlags, v5 interface{}) {
	m.Called(v2, v3, v4, v5)
}
