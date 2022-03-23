package gtk

import (
	"github.com/coyim/gotk3mocks/gdk"
)

type MockAllocation struct {
	gdk.MockRectangle
}

func (m *MockAllocation) GetY() int {
	return m.Called().Int(0)
}
