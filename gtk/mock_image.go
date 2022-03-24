package gtk

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/gtki"
)

type MockImage struct {
	MockWidget
}

func (m *MockImage) SetFromIconName(v1 string, v2 gtki.IconSize) {
	m.Called(v1, v2)
}

func (m *MockImage) Clear() {
	m.Called()
}

func (m *MockImage) SetFromPixbuf(v1 gdki.Pixbuf) {
	m.Called(v1)
}
