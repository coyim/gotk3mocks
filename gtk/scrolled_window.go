package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockScrolledWindow struct {
	MockBin
}

func (m *MockScrolledWindow) GetVAdjustment() gtki.Adjustment {
	return ret[gtki.Adjustment](m.Called(), 0)
}
