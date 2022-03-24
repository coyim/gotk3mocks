package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockMenuShell struct {
	MockContainer
}

func (m *MockMenuShell) Append(v1 gtki.MenuItem) {
	m.Called(v1)
}
