package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockContainer struct {
	MockWidget
}

func (m *MockContainer) Add(v1 gtki.Widget) {
	m.Called(v1)
}

func (m *MockContainer) Remove(v1 gtki.Widget) {
	m.Called(v1)
}

func (m *MockContainer) SetBorderWidth(v1 uint) {
	m.Called(v1)
}

func (m *MockContainer) GetChildren() []gtki.Widget {
	return ret[[]gtki.Widget](m.Called(), 0)
}
