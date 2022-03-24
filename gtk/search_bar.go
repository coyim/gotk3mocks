package gtk

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/gtki"
)

type MockSearchBar struct {
	MockBin
}

func (m *MockSearchBar) ConnectEntry(v1 gtki.Entry) {
	m.Called(v1)
}

func (m *MockSearchBar) GetSearchMode() bool {
	return m.Called().Bool(0)
}

func (m *MockSearchBar) SetSearchMode(v1 bool) {
	m.Called(v1)
}

func (m *MockSearchBar) GetShowCloseButton() bool {
	return m.Called().Bool(0)
}

func (m *MockSearchBar) SetShowCloseButton(v1 bool) {
	m.Called(v1)
}

func (m *MockSearchBar) HandleEvent(v1 gdki.Event) {
	m.Called(v1)
}
