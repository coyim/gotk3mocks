package gdk

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockPixbufLoader struct {
	glib.MockObject
}

func (m *MockPixbufLoader) Close() error {
	return m.Called().Error(0)
}

func (m *MockPixbufLoader) GetPixbuf() (gdki.Pixbuf, error) {
	args := m.Called()
	return ret[gdki.Pixbuf](args, 0), args.Error(1)
}

func (m *MockPixbufLoader) SetSize(v1 int, v2 int) {
	m.Called(v1, v2)
}

func (m *MockPixbufLoader) Write(b []byte) (int, error) {
	args := m.Called(b)
	return args.Int(0), args.Error(1)
}
