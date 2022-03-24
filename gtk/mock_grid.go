package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockGrid struct {
	MockContainer
}

func (m *MockGrid) Attach(v1 gtki.Widget, v2, v3, v4, v5 int) {
	m.Called(v1, v2, v3, v4, v5)
}
