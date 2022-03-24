package gtk

import (
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/coyim/gotk3mocks/glib"
)

type MockBuilder struct {
	glib.MockObject
}

func (m *MockBuilder) AddFromString(v1 string) error {
	return m.Called(v1).Error(0)
}

func (m *MockBuilder) AddFromResource(v1 string) error {
	return m.Called(v1).Error(0)
}

func (m *MockBuilder) GetObject(v1 string) (glibi.Object, error) {
	args := m.Called(v1)
	return ret[glibi.Object](args, 0), args.Error(1)
}

func (m *MockBuilder) ConnectSignals(v1 map[string]interface{}) {
	m.Called(v1)
}
