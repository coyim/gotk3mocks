package gdk

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

type getter interface {
	Get(int) interface{}
}

func ret[T any](args getter, index int) T {
	ret, _ := args.Get(index).(T)
	return ret
}

func (m *Mock) EventButtonFrom(ev gdki.Event) gdki.EventButton {
	return ret[gdki.EventButton](m.Called(ev), 0)
}

func (m *Mock) EventKeyFrom(ev gdki.Event) gdki.EventKey {
	return ret[gdki.EventKey](m.Called(ev), 0)
}

func (m *Mock) PixbufLoaderNew() (gdki.PixbufLoader, error) {
	args := m.Called()
	return ret[gdki.PixbufLoader](args, 0), args.Error(1)
}

func (m *Mock) ScreenGetDefault() (gdki.Screen, error) {
	args := m.Called()
	return ret[gdki.Screen](args, 0), args.Error(1)
}

func (m *Mock) WorkspaceControlSupported() bool {
	return m.Called().Bool(0)
}

func (m *Mock) NewRGBA(values ...float64) gdki.Rgba {
	return ret[gdki.Rgba](m.Called(values), 0)
}
