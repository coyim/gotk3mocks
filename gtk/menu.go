package gtk

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/gtki"
)

type MockMenu struct {
	MockMenuShell
}

func (m *MockMenu) PopupAtMouseCursor(v1 gtki.Menu, v2 gtki.MenuItem, v3 int, v4 uint32) {
	m.Called(v1, v2, v3, v4)
}

func (m *MockMenu) PopupAtPointer(v1 gdki.Event) {
	m.Called(v1)
}
