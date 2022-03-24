package gtk

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockStyleContext struct {
	glib.MockObject
}

func (m *MockStyleContext) AddClass(v1 string) {
	m.Called(v1)
}

func (m *MockStyleContext) RemoveClass(v1 string) {
	m.Called(v1)
}

func (m *MockStyleContext) AddProvider(v1 gtki.StyleProvider, v2 uint) {
	m.Called(v1, v2)
}

func (m *MockStyleContext) GetScreen() (gdki.Screen, error) {
	args := m.Called()
	return ret[gdki.Screen](args, 0), args.Error(1)
}

func (m *MockStyleContext) GetProperty2(v1 string, v2 gtki.StateFlags) (interface{}, error) {
	args := m.Called(v1, v2)
	return args.Get(0), args.Error(1)
}
