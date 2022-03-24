package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockBin struct {
	MockContainer
}

func (m *MockBin) GetChild() gtki.Widget {
	return ret[gtki.Widget](m.Called(), 0)
}
