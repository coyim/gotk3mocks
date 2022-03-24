package gtk

import "github.com/coyim/gotk3mocks/glib"

type MockCssProvider struct {
	glib.MockObject
}

func (m *MockCssProvider) LoadFromData(v1 string) error {
	return m.Called(v1).Error(0)
}
