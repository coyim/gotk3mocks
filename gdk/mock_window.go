package gdk

import "github.com/coyim/gotk3mocks/glib"

type MockWindow struct {
	glib.MockObject
}

func (m *MockWindow) GetDesktop() uint32 {
	return ret[uint32](m.Called(), 0)
}

func (m *MockWindow) MoveToDesktop(v1 uint32) {
	m.Called(v1)
}
